package controllers

import (
	temp "groupie-tracker/templates"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "home", nil)
}
