package controllers

import (
	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
	"golang-boilerplate/utils"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := models.GetUserByEmail2(db, email)
		if err != nil {
			log.Println("Invalid email or password")
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			log.Println("Invalid email or password")
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		// Successful login, redirect to dashboard
		http.Redirect(w, r, "/backend/dashboard", http.StatusSeeOther)
		return
	}

	// If not a POST request, show the login page
	ShowLoginPage(w, r)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirmation := r.FormValue("password_confirmation")

		// Verify and process registration
		if name != "" && username != "" && email != "" && password != "" && password == passwordConfirmation {
			hashedPassword, err := utils.HashPassword(password)
			if err != nil {
				log.Printf("Error hashing password: %v", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			// Create user in database
			err = models.CreateUser(db, username, name, email, hashedPassword)
			if err != nil {
				log.Printf("Error creating user: %v", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			// Redirect to login page on successful registration
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		// If registration fails, redirect to register page
		http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
		return
	}

	// If not a POST request, show the register page
	ShowRegisterPage(w, r)
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
