package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Search_Request(query string, page, perPage int) (HomePagePhotos, int, error) {
	// URL encode the search query
	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&page=%d&per_page=%d",
		encodedQuery, page, perPage)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError,
			fmt.Errorf("Error initializing request: %v", reqErr)
	}

	req.Header.Set("Authorization", _clientApiKey)

	res, resErr := _httpClient.Do(req)
	if resErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError,
			fmt.Errorf("Error sending request: %v", resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return HomePagePhotos{}, res.StatusCode,
			fmt.Errorf("Error in response: code %d", res.StatusCode)
	}

	var data HomePagePhotos
	decodeErr := json.NewDecoder(res.Body).Decode(&data)
	if decodeErr != nil {
		return HomePagePhotos{}, http.StatusInternalServerError,
			fmt.Errorf("Error decoding data: %v", decodeErr)
	}

	return data, res.StatusCode, nil
}
