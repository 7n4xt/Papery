package routes

import (
    "net/http"
	"groupie-tracker/controllers"
)

func FavoriteRoutes() {
    http.HandleFunc("/favorite", controllers.FavoritePage)
    http.HandleFunc("/api/favorites/add", controllers.AddFavorite)
    http.HandleFunc("/api/favorites/remove", controllers.RemoveFavorite)
}