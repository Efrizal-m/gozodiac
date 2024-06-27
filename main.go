package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Struct untuk menyimpan data form
type FormData struct {
	Name      string
	BirthDate string
}

// Struct untuk menyimpan hasil
type ResultData struct {
	Name      string
	AgeYears  int
	AgeMonths int
	AgeDays   int
	Zodiac    string
}

// Fungsi untuk menghitung usia
func calculateAge(birthDate time.Time) (years, months, days int) {
	today := time.Now()
	years = today.Year() - birthDate.Year()
	months = int(today.Month()) - int(birthDate.Month())
	days = today.Day() - birthDate.Day()

	if days < 0 {
		months--
		days += 30 // simplifikasi
	}
	if months < 0 {
		years--
		months += 12
	}
	return
}

// Fungsi untuk mendapatkan zodiak
func getZodiac(db *sql.DB, birthDate time.Time) string {
	var zodiac string
	dateStr := birthDate.Format("02-Jan")
	query := "SELECT ZodiacName FROM TZodiac WHERE (StartDate <= ? AND EndDate >= ?) OR (StartDate > EndDate AND (StartDate <= ? OR EndDate >= ?))"
	err := db.QueryRow(query, dateStr, dateStr, dateStr, dateStr).Scan(&zodiac)
	if err != nil {
		log.Fatal(err)
	}
	return zodiac
}

// Fungsi untuk menampilkan form
func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

// Fungsi untuk memproses form
func submitHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./zodiac.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		birthDateStr := r.FormValue("birthdate")
		// Check if birthDateStr is empty
		if birthDateStr == "" {
			http.Error(w, "Birthdate cannot be empty", http.StatusBadRequest)
			return
		}

		birthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			log.Fatal(err)
		}

		years, months, days := calculateAge(birthDate)
		zodiac := getZodiac(db, birthDate)

		result := ResultData{
			Name:      name,
			AgeYears:  years,
			AgeMonths: months,
			AgeDays:   days,
			Zodiac:    zodiac,
		}

		tmpl := template.Must(template.ParseFiles("result.html"))
		tmpl.Execute(w, result)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", formHandler).Methods("GET")
	r.HandleFunc("/submit", submitHandler).Methods("POST")

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
