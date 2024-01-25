package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatalf("[main] database connection error: %v", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute)

	count := 10

	for i := 0; i < count; i++ {
		email := gofakeit.Email()
		password := gofakeit.Password(true, true, true, false, false, 10)
		name := gofakeit.NamePrefix()
		lastName := gofakeit.LastName()
		districtID := gofakeit.Number(1, 4)

		query := `
			INSERT INTO user (email, username, password, name, lastName, districtID)
			VALUES (?, ?, ?, ?, ?, ?)`
		_, err := db.Exec(query, email, email, password, name, lastName, districtID)

		if err != nil {
			log.Fatalf("[main] database insert error: %v", err.Error())
		}

		log.Printf("[main] %v. user created", i+1)

	}
}
