package models

import (
	"database/sql"
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

func CreateUser(db *sql.DB, name, email, password string) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	return err
}
