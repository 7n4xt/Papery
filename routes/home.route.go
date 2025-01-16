package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}
