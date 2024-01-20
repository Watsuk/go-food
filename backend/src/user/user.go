package user

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(db *sql.DB) ([]entity.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []entity.User
	var createdAt []uint8
	var updatedAt []uint8
	var deletedAt []uint8
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Fatal(err)
		}
		createdAtString := string(createdAt)

		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			log.Fatal(err)
		}

		updatedAtString := string(updatedAt)
		user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtString)
		if err != nil {
			log.Fatal(err)
		}
		if deletedAt != nil {
			deletedAtString := string(deletedAt)
			user.DeletedAt, err = time.Parse("2006-01-02 15:04:05", deletedAtString)
			if err != nil {
				log.Fatal(err)
			}
		}

		users = append(users, user)
	}

	return users, err
}

func CreateUser(db *sql.DB, userName string, password string, email string, role int) (entity.User, error) {

	user := entity.User{
		Username: userName,
		Password: password,
		Email:    email,
		Role:     role,
	}

	_, err := db.Exec("INSERT INTO users (username, pw_hash, email, permissions) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.Email, user.Role)
	if err != nil {
		log.Printf("Erreur lors de la création de l'utilisateur : %v", err)
		return entity.User{}, fmt.Errorf("could not create user: %v", err)
	}

	return user, err
}
