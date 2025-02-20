package routes

import (
	"groupie-tracker/controllers"
	"net/http"
)

func FavoriteRoutes() {
	http.HandleFunc("/favorite/", controllers.FavoritePage)
}