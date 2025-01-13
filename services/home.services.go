package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Déclaration du client qui va émettre les requêtes
var _httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// la structure de la page d'accueil
type PhotoSrc struct {
	Original string `json:"original"`
	Large2x  string `json:"large2x"`
}

type Photo struct {
	Src PhotoSrc `json:"src"`
}

type HomepagePhotos struct {
	Page    int     `json:"page"`
	PerPage int     `json:"per_page"`
	Photos  []Photo `json:"photos"`
}

// Initialisation des variables pour s'authentifier à l'API
var _clientApiKey string = "Y2W6tV0zwZjAUd84QZDkUOPuviZaXHGxuShzBuvbxstGnHjBzgb5X8pI"

// HomePagePhotosRequest fetches photos from the Pexels API.
func HomePagePhotosRequest(query string) (HomepagePhotos, int, error) {
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s", query)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return HomepagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors de l'initialisation de la requête : %v", reqErr)
	}

	req.Header.Set("Authorization", _clientApiKey)

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return HomepagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors de l'envoi de la requête : %v", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return HomepagePhotos{}, res.StatusCode, fmt.Errorf("Erreur dans la réponse : code %d", res.StatusCode)
	}

	var data HomepagePhotos
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return HomepagePhotos{}, http.StatusInternalServerError, fmt.Errorf("Erreur lors du décodage des données : %v", decodeErr)
	}

	return data, res.StatusCode, nil
}
