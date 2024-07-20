package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

// SetSessionStore sets the session store for the middleware
func SetSessionStore(s *sessions.CookieStore) {
	store = s
}

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/backend/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
