package controllers

import (
	//"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func DetailsPage(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "details", nil)
}
