package controllers

import (
	"groupie-tracker/services"
	temp "groupie-tracker/templates"
	"net/http"
	"strconv"
)

type SearchPageData struct {
	Photos      []services.Photo
	Query       string
	CurrentPage int
	HasNextPage bool
	HasPrevPage bool
	NextPage    int
	PrevPage    int
	Filter string
}

// Update search.controller.go
func SearchPage(w http.ResponseWriter, r *http.Request) {
    // Get the search query from URL parameters
    query := r.URL.Query().Get("search-word")
    if query == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }


    // Get the page parameter, defaulting to 1 if not present
    pageStr := r.URL.Query().Get("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }

    // Set the number of images per page
    perPage := 37

    // Fetch photos based on search query with pagination
    data, status, err := services.Search_Request(query, page, perPage)
    if err != nil {
        http.Error(w, err.Error(), status)
        return
    }

    // Prepare data for the template
    pageData := SearchPageData{
        Photos:      data.Photos,
        Query:       query,
        CurrentPage: page,
        HasNextPage: (page * perPage) < data.TotalResults,
        HasPrevPage: page > 1,
        NextPage:    page + 1,
        PrevPage:    page - 1,
    }
	
    err = temp.Temp.ExecuteTemplate(w, "search", pageData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}