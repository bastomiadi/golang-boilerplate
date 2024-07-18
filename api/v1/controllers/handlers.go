// api/v1/controllers/handlers.go
package controllers

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API v1 module")
}
