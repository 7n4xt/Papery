package main

import (
    "groupie-tracker/routes"
    "groupie-tracker/services"
    temp "groupie-tracker/templates"
    "log"
)

func main() {
    // Initialize favorites system
    if err := services.InitFavorites(); err != nil {
        log.Fatal("Failed to initialize favorites:", err)
    }
    temp.InitTemplates()
    routes.InitServe()
}