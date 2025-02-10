package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func DetailsRoutes() {
	http.HandleFunc("/details/", controllers.DetailsPage)
}
