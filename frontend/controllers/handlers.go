// frontend/controllers/handlers.go
package controllers

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Frontend module")
}
