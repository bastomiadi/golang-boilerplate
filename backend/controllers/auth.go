package controllers

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/utils"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "auth")

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := models.GetUserByEmail(email)
		if err != nil {
			log.Println("Invalid email or password")
			session.Values["login_error"] = "Invalid email or password"
			session.Save(r, w)
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		// Use the CheckPasswordHash function from the utils package
		if !utils.CheckPasswordHash(password, user.Password) {
			log.Println("Invalid email or password")
			session.Values["login_error"] = "Invalid email or password"
			session.Save(r, w)
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}

		// Successful login, redirect to dashboard
		session.Values["authenticated"] = true
		delete(session.Values, "login_error")
		session.Save(r, w)
		http.Redirect(w, r, "/backend/dashboard", http.StatusSeeOther)
		return
	}

	// If not a POST request, show the login page
	ShowLoginPage(w, r)
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "auth")

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirmation := r.FormValue("password_confirmation")
		terms := r.FormValue("terms")

		if name == "" || username == "" || email == "" || password == "" || password != passwordConfirmation {
			session.Values["register_error"] = "All fields are required and passwords must match."
			session.Save(r, w)
			http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
			return
		}

		if terms != "agree" {
			session.Values["register_error"] = "You must agree to the terms."
			session.Save(r, w)
			http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
			return
		}

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			session.Values["register_error"] = "Server error. Please try again later."
			session.Save(r, w)
			http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
			return
		}

		err = models.CreateUser(username, name, email, hashedPassword)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			session.Values["register_error"] = "Server error. Please try again later."
			session.Save(r, w)
			http.Redirect(w, r, "/backend/register", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
		return
	}

	ShowRegisterPage(w, r)
}

func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")
	loginError := ""
	if err, ok := session.Values["login_error"].(string); ok {
		loginError = err
		delete(session.Values, "login_error")
		session.Save(r, w)
	}

	tmpl, err := template.ParseFiles("backend/views/auth/login.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	data := struct {
		LoginError string
	}{
		LoginError: loginError,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")
	registerError := ""
	if err, ok := session.Values["register_error"].(string); ok {
		registerError = err
		delete(session.Values, "register_error")
		session.Save(r, w)
	}

	tmpl, err := template.ParseFiles("backend/views/auth/register.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	data := struct {
		RegisterError string
	}{
		RegisterError: registerError,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")
	// Revoke user authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	// Redirect to login page
	http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
}
