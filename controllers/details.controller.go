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
    
    // Get filter from query params, default to "large"
    filter := r.URL.Query().Get("filter")
    if filter == "" {
        filter = "large"
    }

    // Get photo details from service
    photo, statusCode, err := services.GetPhotoDetails(photoID)
    if err != nil || statusCode != http.StatusOK {
        http.Error(w, "Error retrieving photo details", statusCode)
        return
    }
    
    // Create template data
    data := struct {
        Photo  *services.Photo
        Filter string
    }{
        Photo:  photo,
        Filter: filter,
    }

    temp.Temp.ExecuteTemplate(w, "details", data)
}
