package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPhotoDetails(photoID string) (*Photo, int, error) {
	url := fmt.Sprintf("https://api.pexels.com/v1/photos/%s", photoID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Authorization", _clientApiKey)

	res, err := _httpClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("API returned status code %d", res.StatusCode)
	}

	var photo Photo
	if err := json.NewDecoder(res.Body).Decode(&photo); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &photo, http.StatusOK, nil
}
