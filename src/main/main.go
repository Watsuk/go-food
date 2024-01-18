package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func main() {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "buffet",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := database.CreateStore(db)
	mux := web.CreateMux(store)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		_ := fmt.Errorf("could not listen on port 8080 %v", err)
		return
	}

}
