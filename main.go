package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AhmetSBulbul/quarterback-server/gapi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatalf("[main] database connection error: %v", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute)

	go func() {
		r := mux.NewRouter()
		r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	gapi.NewServer(db)
}
