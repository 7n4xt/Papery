package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func AboutRoutes() {
	http.HandleFunc("/about", controllers.AboutPage)
}
