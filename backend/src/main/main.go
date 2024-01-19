package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/handler"
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

	user := &entity.User{
		ID:        1,
		Username:  "watsuk",
		Password:  "1234",
		Email:     "test@gmail.com",
		Role:      1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mux := handler.NewHandlerUser(user)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Errorf("could not listen on port 8080 %v", err)
		return
	}

}
