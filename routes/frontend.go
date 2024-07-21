package routes

import (
	frontendControllers "golang-boilerplate/frontend/controllers"

	"github.com/gorilla/mux"
)

// Add your frontend controllers
func RegisterFrontendRoutes(router *mux.Router) {
	// Define frontend routes here
	router.HandleFunc("/frontend/", frontendControllers.HandleRequest)

	// router.HandleFunc("/", frontendControllers.HandleRequest)
}
