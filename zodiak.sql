-- Buat database dengan SQLite3
CREATE TABLE TZodiac (
    StartDate TEXT,
    EndDate TEXT,
    ZodiacName TEXT
);

-- Isi data zodiak
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('22-Dec', '20-Jan', 'Capricorn');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Jan', '19-Feb', 'Aquarius');
INSERT INTO TZodiac (StartDate, EndDate, ZodiacName) VALUES ('21-Mar', '20-Apr', 'Aries');
-- Tambahkan semua data zodiak lainnya
