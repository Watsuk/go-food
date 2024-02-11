package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/handler"
	"github.com/go-sql-driver/mysql"
)

func main() {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "db",
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
	var reference entity.Reference
	reference.User = &entity.User{}
	reference.Truck = &entity.Truck{}
	reference.Product = &entity.Product{}
	reference.Order = &entity.Order{}

	mux := handler.NewHandler(db, reference)

	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("could not listen on port 3000: %v", err)
		return
	}

}
