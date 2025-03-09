package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Search_Request(query string, page, perPage int, filters map[string]string) (HomePagePhotos, int, error) {
	// URL encode the search query
	encodedQuery := url.QueryEscape(query)

	// Start building the URL with required parameters
	baseURL := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&page=%d&per_page=%d",
		encodedQuery, page, perPage)

	// Add optional filters if they are present and not empty
	params := url.Values{}
	if orientation, ok := filters["orientation"]; ok && orientation != "" {
		params.Add("orientation", orientation)
	}
	if size, ok := filters["size"]; ok && size != "" {
		params.Add("size", size)
	}
	if color, ok := filters["color"]; ok && color != "" {
		params.Add("color", color)
	}

	// Append filters to base URL if any exist
	finalURL := baseURL
	if len(params) > 0 {
		finalURL += "&" + params.Encode()
	}

	req, reqErr := http.NewRequest(http.MethodGet, finalURL, nil)
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