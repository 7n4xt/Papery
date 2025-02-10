package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}

func searchRoutes() {
	http.HandleFunc("/search-handle/", controllers.SearchPage)
}
