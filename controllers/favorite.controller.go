package controllers

import (
	temp "groupie-tracker/templates"
	"net/http"
)

func FavoritePage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "details", nil)
}
