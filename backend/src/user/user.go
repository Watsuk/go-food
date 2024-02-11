package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Watsuk/go-food/src/entity"
	"github.com/Watsuk/go-food/src/permissions"
	"github.com/Watsuk/go-food/src/tokens"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(db *sql.DB) ([]entity.User, error) {
	rows, err := db.Query("SELECT id, username, email, permissions, created_at, updated_at, deleted_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []entity.User
	var createdAt, updatedAt, deleted_at []uint8
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &createdAt, &updatedAt, &deleted_at)
		if err != nil {
			log.Fatal(err)
		}
		if deleted_at != nil {
			continue
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

		users = append(users, user)
	}

	return users, err
}

func GetUserByID(db *sql.DB, userID int) (*entity.User, error) {
	row := db.QueryRow("SELECT id, username, email, permissions, created_at, updated_at, deleted_at FROM users WHERE id = ?", userID)

	var user entity.User
	var createdAt, updatedAt, deletedAt []uint8

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatal(err)
	}
	if deletedAt != nil {
		return nil, nil
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

	return &user, nil
}

func CreateUser(db *sql.DB, userName string, password string, email string, role int) (entity.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erreur lors du hashage du mot de passe : %v", err)
		return entity.User{}, fmt.Errorf("could not hash password: %v", err)
	}

	user := entity.User{
		Username: userName,
		Password: string(hashedPassword),
		Email:    email,
		Role:     permissions.Permission(role),
	}

	_, err = db.Exec("INSERT INTO users (username, pw_hash, email, permissions) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.Email, int(user.Role))
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
			return 0, "", errors.New("invalid email")

		}
		return 0, "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return 0, "", errors.New("invalid password")
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, tokens.Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	token, err := tk.SignedString([]byte("SecretKey"))
	if err != nil {
		return 0, "", err
	}

	// tokenString, err := token.SignedString([]byte("SecretKey"))
	// if err != nil {
	// 	return 0, "", err
	// }

	return user.ID, token, nil
}

func EditUser(db *sql.DB, userID int64, username string, email string, role int) error {
	_, err := db.Exec("UPDATE users SET username = ?, email = ?, permissions = ? WHERE id = ?", username, email, role, userID)
	return err
}

func DeleteUser(db *sql.DB, userID int64) error {
	_, err := db.Exec("UPDATE users SET deleted_at = ? WHERE id = ?", time.Now(), userID)
	fmt.Println(err)
	return err
}
