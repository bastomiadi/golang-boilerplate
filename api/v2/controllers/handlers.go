// api/v2/controllers/handlers.go
package controllers

import (
	"encoding/json"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Your logic for V2 API
	response := map[string]string{"message": "API V2 Endpoint"}
	json.NewEncoder(w).Encode(response)
}
