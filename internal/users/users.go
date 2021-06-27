package users

import (
	"database/sql"
	"log"

	database "github.com/saputradharma/go-graphql-example/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (user *User) Create() error {
	query := `INSERT INTO users (username, password) VALUES(?, ?)`

	hashedPassword, err := HashPassword(user.Password)
	_, err = database.DB.Exec(query, user.Username, hashedPassword)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (user *User) Authenticate() bool {
	var hashedPassword string
	var row *sql.Row

	query := `SELECT password FROM users WHERE username = ?`

	row = database.DB.QueryRow(query, user.Username)

	err := row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Println(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	var id int
	var row *sql.Row

	query := `SELECT id FROM users WHERE username = ?`

	row = database.DB.QueryRow(query, username)

	err := row.Scan(&id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}
