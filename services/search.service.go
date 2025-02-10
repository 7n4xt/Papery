package services

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/controllers"
	"net/http"
)

func Search_Request() (HomePagePhotos, int, error) {

	var Search string
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&per_page=40", Search)

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
