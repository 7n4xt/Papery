package controllers

import (
	temp "groupie-tracker/templates"
	"net/http"
)

func AboutPage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "details", nil)
}
