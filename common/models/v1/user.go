package models

import (
	"golang-boilerplate/config"
	"golang-boilerplate/utils"
)

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(255);unique;not null" json:"username"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"` // Exclude from database and JSON serialization
}

// CreateUser creates a new user in the database
func CreateUser(name, username, email, password string) error {
	db := config.GetDB()

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return db.Create(&user).Error
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (*User, error) {
	db := config.GetDB()

	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email from the database
func GetUserByEmail(email string) (*User, error) {
	db := config.GetDB()

	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
