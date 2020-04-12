package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type dbconnection struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func main() {
	var dbcon dbconnection
	dbcon.host = os.Getenv("DB_HOST")
	dbcon.password = os.Getenv("DB_PASSWORD")
	dbcon.user = os.Getenv("DB_USER")
	dbcon.dbname = os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if port == "" {
		dbcon.port = 5432
	} else {
		dbcon.port, _ = strconv.Atoi(port)
	}
	if dbcon.host == "" {
		dbcon.host = "localhost"
	}
	fmt.Printf("%+v\n", dbcon)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := connectToDB(dbcon)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3089", nil)
}

func connectToDB(dbcon dbconnection) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbcon.host, dbcon.port, dbcon.user, dbcon.password, dbcon.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")
	return nil
}
