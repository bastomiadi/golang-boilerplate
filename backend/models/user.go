package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func CreateUser(db *sql.DB, name, username, email, password string) error {
	_, err := db.Exec("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)", name, username, email, password)
	return err
}

func (u *User) Create(db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, name, email, password) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, u.Username, u.Name, u.Email, string(hashedPassword))
	return err
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, name, email, password FROM users WHERE username = ?`
	row := db.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail retrieves a user by email from the database
func GetUserByEmail2(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, name, email, password FROM users WHERE email = ?`
	row := db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
