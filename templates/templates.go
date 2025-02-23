package temp

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
)

var Temp *template.Template

func InitTemplates() {
    // Serve static files
    fileServer := http.FileServer(http.Dir("./assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
    
    // Parse templates
    temp, tempErr := template.ParseGlob("./templates/*.html")
    if tempErr != nil {
        fmt.Printf("Erreur Template - Une erreur lors du chargement des template \n message d'erreur : %v\n", tempErr.Error())
        os.Exit(1)
    }
    Temp = temp
}