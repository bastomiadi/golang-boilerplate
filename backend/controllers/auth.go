package controllers

import (
	"database/sql"
	"golang-boilerplate/backend/models"
	"golang-boilerplate/utils"
	"html/template"
	"log"
	"net/http"
)

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Example login logic
		// Verify email and password here, this is a simplified example
		if email == "test@example.com" && password == "password" {
			// Successful login, redirect to dashboard
			http.Redirect(w, r, "/backend/dashboard", http.StatusSeeOther)
			return
		}

		// If login fails, you can show an error message or redirect to login page
		http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
		return
	}

	// If not a POST request, show the login page
	ShowLoginPage(w, r)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirmation := r.FormValue("password_confirmation")

		// Example registration logic
		// Verify and process registration, this is a simplified example
		if name != "" && email != "" && password != "" && password == passwordConfirmation {
			// Example of hashing password before storing (use your utility function)
			hashedPassword, err := utils.HashPassword(password)
			if err != nil {
				log.Printf("Error hashing password: %v", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			// Create user in database
			db := getDB()
			err = models.CreateUser(db, name, email, hashedPassword)
			if err != nil {
				log.Printf("Error creating user: %v", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			// Redirect to login page on successful registration
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		// If registration fails, you can show an error message or redirect to register page
		http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
		return
	}

	// If not a POST request, show the register page
	ShowRegisterPage(w, r)
}

func getDB() *sql.DB {
	// Example function to get database connection, replace with your actual implementation
	db, err := sql.Open("mysql", "root:@/golang-boilerplate")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	return db
}

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("backend/views/auth/login.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("backend/views/auth/register.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}
