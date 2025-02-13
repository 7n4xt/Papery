package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// HTTP client with timeout
var _httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// API response structure
type PhotoSrc struct {
	Original string `json:"original"`
	Large2x  string `json:"large2x"`
	Large    string `json:"large"`
}

type Photo struct {
	ID              int      `json:"id"`
	Src             PhotoSrc `json:"src"`
	Alt             string   `json:"alt"`
	Photographer    string   `json:"photographer"`
	PhotoGrapherURL string   `json:"photographer_url"`
}

type HomePagePhotos struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	Photos       []Photo `json:"photos"`
}

// API Key (Replace this with a secure method to load your API key)
const _clientApiKey string = "Y2W6tV0zwZjAUd84QZDkUOPuviZaXHGxuShzBuvbxstGnHjBzgb5X8pI"

// Fetch photos with pagination
func HomePagePhotosRequest(page, perPage int) (HomePagePhotos, int, error) {
	var randomeQuery = []string{"wallpapers", "abstract", "nature", "Cars", "technology"}

	url := fmt.Sprintf("https://api.pexels.com/v1/search?%s=wallpapers&page=%d&per_page=%d", randomeQuery[rand.Int(4)], page, perPage)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors de l'initialisation de la requête : %v", reqErr)
	}

	req.Header.Set("Authorization", _clientApiKey)

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors de l'envoi de la requête : %v", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return HomePagePhotos{}, res.StatusCode, fmt.Errorf("Erreur dans la réponse : code %d", res.StatusCode)
	}

	var data HomePagePhotos
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors du décodage des données : %v", decodeErr)
	}

	return data, res.StatusCode, nil
}
