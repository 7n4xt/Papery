package controllers

import (
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	data, statusCode, err := services.HomePagePhotosRequest()
	if err != nil || statusCode != http.StatusOK {
		http.Error(w, "erreur lors de la recuperation des photos", statusCode)
		return
	}
	temp.Temp.ExecuteTemplate(w, "home", data)
}
