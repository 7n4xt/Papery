package routes

import (
    "fmt"
    "log"
    "net/http"
)

func InitServe() {
    FavoriteRoutes()
    homeRoutes()
    DetailsRoutes()
    searchRoutes()
    
    fmt.Println("Le serveur est opérationel : http://localhost:8080")
    if err := http.ListenAndServe("localhost:8080", nil); err != nil {
        log.Fatal("Server error:", err)
    }
}