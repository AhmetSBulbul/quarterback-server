package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/AhmetSBulbul/quarterback-server/helpers"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func users(db *sql.DB) {
	count := 10
	for i := 0; i < count; i++ {
		email := gofakeit.Email()
		password := helpers.HashPassword(gofakeit.Password(true, true, true, false, false, 10))
		name := gofakeit.FirstName()
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

func games(db *sql.DB) {
	count := 10

	query := `
		INSERT INTO game (courtID, startedAt, endedAt, homeScore, awayScore)
		VALUES (?, ?, ?, ?, ?);
	`

	for i := 0; i < count; i++ {
		courtID := gofakeit.Number(1, 9)
		startedAt := gofakeit.PastDate()
		endedAt := gofakeit.DateRange(startedAt, time.Now())
		homeScore := gofakeit.Number(0, 100)
		awayScore := gofakeit.Number(0, 100)
		_, err := db.Exec(query, courtID, startedAt, endedAt, homeScore, awayScore)
		if err != nil {
			log.Fatalf("[main] database insert error: %v", err.Error())
		}
		log.Printf("[main] %v. game created", i+1)
	}

	count = 40
	query = `INSERT INTO game_player (gameID, playerID, isHomeSide, isCanceled) VALUES (?, ?, ?, ?);`

	for i := 0; i < count; i++ {
		gameID := gofakeit.Number(1, 10)
		playerID := gofakeit.Number(1, 10)
		isHomeSide := gofakeit.Bool()
		isCanceled := gofakeit.Bool()
		_, err := db.Exec(query, gameID, playerID, isHomeSide, isCanceled)
		if err != nil {
			log.Fatalf("[main] database insert error: %v", err.Error())
		}
		log.Printf("[main] %v. game_player created", i+1)
	}

}

func main() {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatalf("[main] database connection error: %v", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute)

	users(db)
	games(db)
}
