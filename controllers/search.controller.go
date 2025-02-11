package controllers

import (
	"groupie-tracker/services"
	"html/template"
	"net/http"
)

type SearchPageData struct {
	Photos []services.Photo
	Query  string
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
	// Get the search query from URL parameters
	query := r.URL.Query().Get("search-word")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Fetch photos based on search query
	data, status, err := services.Search_Request(query)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	// Prepare data for the template
	pageData := SearchPageData{
		Photos: data.Photos,
		Query:  query,
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "search", pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
