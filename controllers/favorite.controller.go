package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"io"
	"net/http"
	"strconv"
)

func FavoritePage(w http.ResponseWriter, r *http.Request) {
	favorites, err := services.LoadFavorites()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading favorites: %v", err), http.StatusInternalServerError)
		return
	}

	if err := temp.Temp.ExecuteTemplate(w, "favorite", favorites); err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	// Log the received data for debugging
	fmt.Printf("Received data: %s\n", string(body))

	var photo services.Photo
	if err := json.Unmarshal(body, &photo); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if photo.ID == 0 || photo.Src.Original == "" || photo.Src.Large == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if err := services.AddToFavorites(photo); err != nil {
		http.Error(w, fmt.Sprintf("Error adding to favorites: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	photoID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(photoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid photo ID: %v", err), http.StatusBadRequest)
		return
	}

	if err := services.RemoveFromFavorites(id); err != nil {
		http.Error(w, fmt.Sprintf("Error removing from favorites: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
