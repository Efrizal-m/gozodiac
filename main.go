package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Zodiac struct {
	StartDate  string
	EndDate    string
	ZodiacName string
}

type Result struct {
	Name      string
	AgeYears  int
	AgeMonths int
	AgeDays   int
	Zodiac    string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./zodiac.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	dob := r.FormValue("dob")

	birthdate, err := time.Parse("2006-01-02", dob)
	if err != nil {
		log.Fatal(err)
	}

	ageYears, ageMonths, ageDays := calculateAge(birthdate)
	zodiac := getZodiac(birthdate)

	result := Result{
		Name:      name,
		AgeYears:  ageYears,
		AgeMonths: ageMonths,
		AgeDays:   ageDays,
		Zodiac:    zodiac,
	}

	tmpl := template.Must(template.ParseFiles("result.html"))
	tmpl.Execute(w, result)
}

func calculateAge(birthdate time.Time) (years int, months int, days int) {
	now := time.Now()
	years = now.Year() - birthdate.Year()
	months = int(now.Month() - birthdate.Month())
	days = now.Day() - birthdate.Day()

	if days < 0 {
		months--
		days += 30
	}

	if months < 0 {
		years--
		months += 12
	}

	return years, months, days
}

func getZodiac(birthdate time.Time) string {
	rows, err := db.Query("SELECT StartDate, EndDate, ZodiacName FROM TZodiac")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var startDate, endDate string
	var zodiacName string

	for rows.Next() {
		err := rows.Scan(&startDate, &endDate, &zodiacName)
		if err != nil {
			log.Fatal(err)
		}

		start, err := time.Parse("02-Jan", startDate)
		if err != nil {
			log.Fatal(err)
		}
		end, err := time.Parse("02-Jan", endDate)
		if err != nil {
			log.Fatal(err)
		}

		if (birthdate.Month() == start.Month() && birthdate.Day() >= start.Day()) ||
			(birthdate.Month() == end.Month() && birthdate.Day() <= end.Day()) {
			return zodiacName
		}
	}
	return ""
}
