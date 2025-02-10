package controllers

import (
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
	"strings"
)

func DetailsPage(w http.ResponseWriter, r *http.Request) {
	// Extract photo ID from URL
	photoID := strings.TrimPrefix(r.URL.Path, "/details/")

	// Get photo details from service
	photo, statusCode, err := services.GetPhotoDetails(photoID)
	if err != nil || statusCode != http.StatusOK {
		http.Error(w, "Error retrieving photo details", statusCode)
		return
	}

	temp.Temp.ExecuteTemplate(w, "details", photo)
}
