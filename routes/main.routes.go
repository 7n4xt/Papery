package routes

import (
	"fmt"
	"net/http"
)

func InitServe() {

	homeRoutes()
	DetailsRoutes()
	fmt.Println("Le serveur est opérationel : http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
