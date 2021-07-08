package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/db-connect", func(w http.ResponseWriter, req *http.Request) {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("connection refused to the DB: %v \nconnection info: %s", err, psqlInfo)))
			return
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("connection refused to the DB: %v \nconnection info: %s", err, psqlInfo)))
			return
		}

		w.Write([]byte("connected"))
	})

	log.Println("Ready to accept requests")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
