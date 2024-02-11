package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/dgrijalva/jwt-go"
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

func GetUserByID(db *sql.DB, id int) (*entity.User, error) {
	rows := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var user entity.User

	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatal(err)
	}

	return &user, nil
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

func Login(db *sql.DB, email, password string) (int64, string, error) {
	var user entity.User
	err := db.QueryRow("SELECT id, pw_hash FROM users WHERE email = ?", email).Scan(&user.ID, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", errors.New("invalid email or password")

		}
		return 0, "", err
	}

	if password != user.Password {
		return 0, "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
	})

	tokenString, err := token.SignedString([]byte("SekretKey"))
	if err != nil {
		return 0, "", err
	}

	return user.ID, tokenString, nil
}

func EditUser(db *sql.DB, userID int64, username string, email string, role int) error {
	_, err := db.Exec("UPDATE users SET username = ?, email = ?, permissions = ? WHERE id = ?", username, email, role, userID)
	return err
}

func DeleteUser(db *sql.DB, userID int64) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	return err
}
