package token

import (
	"database/sql"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
)

func GetTokens(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT token FROM tokens")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var tokens []string

	for rows.Next() {
		var token entity.Token
		err = rows.Scan(&token.Token)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tokens = append(tokens, token.Token)
	}

	return tokens, nil
}

func AddToken(db *sql.DB, token string, userID int64) error {
	_, err := db.Exec("INSERT INTO tokens (token, user_id, lifetime) VALUES (?, ?, ?)", token, userID, time.Now().Add(time.Hour*24))
	return err
}
