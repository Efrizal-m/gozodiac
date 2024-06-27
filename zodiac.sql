-- Buat database dengan SQLite3
CREATE TABLE TZodiac (
    StartDate TEXT,
    EndDate TEXT,
    ZodiacName TEXT
);

-- Isi data zodiak
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('22-Dec', '20-Jan', 'Capricorn');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Jan', '19-Feb', 'Aquarius');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('20-Feb', '20-Mar', 'Pisces');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Mar', '20-Apr', 'Aries');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Apr', '20-May', 'Taurus');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-May', '20-Jun', 'Gemini');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Jun', '22-Jul', 'Cancer');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('23-Jul', '22-Aug', 'Leo');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('23-Aug', '22-Sep', 'Virgo');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('23-Sep', '22-Oct', 'Libra');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('23-Oct', '21-Nov', 'Scorpio');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('22-Nov', '21-Dec', 'Sagittarius');