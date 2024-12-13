package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Déclaration du client qui va émettre les requêtes
var _httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// Initialisation des variables pour s'authentifier à l'API
var _clientApiKey string = "Y2W6tV0zwZjAUd84QZDkUOPuviZaXHGxuShzBuvbxstGnHjBzgb5X8pI"

func 