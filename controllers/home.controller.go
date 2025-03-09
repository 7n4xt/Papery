package controllers

import (
    "groupie-tracker/services"
    temp "groupie-tracker/templates"
    "net/http"
    "strconv"
)

// Update home.controller.go
func HomePage(w http.ResponseWriter, r *http.Request) {
    // Get the page number from the form value
    pageStr := r.FormValue("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1 // Default to page 1 if there's an error or invalid input
    }

    // Set the number of images per page
    perPage := 48

    // Fetch the photos from the API with pagination
    data, statusCode, err := services.HomePagePhotosRequest(page, perPage)
    if err != nil || statusCode != http.StatusOK {
        http.Error(w, "Erreur lors de la récupération des photos", statusCode)
        return
    }

    // Pass pagination data to the template
    temp.Temp.ExecuteTemplate(w, "home", map[string]interface{}{
        "Photos":      data.Photos,
        "CurrentPage": page,
        "HasNextPage": (page * perPage) < data.TotalResults,
        "HasPrevPage": page > 1,
        "NextPage":    page + 1,
        "PrevPage":    page - 1,
    })
}