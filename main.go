package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/AhmetSBulbul/quarterback-server/gapi"
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
	gapi.NewServer(db)
}
