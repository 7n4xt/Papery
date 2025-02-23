package services

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "sync"
)

type Favorite struct {
    Photos []Photo `json:"photos"`
}

var (
    favoritesFile = "data/favorites.json"
    mutex         sync.RWMutex
)

// Initialize favorites file if it doesn't exist
func InitFavorites() error {
    // Create data directory if it doesn't exist
    dataDir := filepath.Dir(favoritesFile)
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        return fmt.Errorf("failed to create data directory: %v", err)
    }

    // Check if favorites file exists
    if _, err := os.Stat(favoritesFile); os.IsNotExist(err) {
        // Create new favorites file with empty array
        favorites := Favorite{Photos: []Photo{}}
        data, err := json.MarshalIndent(favorites, "", "    ")
        if err != nil {
            return fmt.Errorf("failed to marshal initial favorites: %v", err)
        }

        if err := os.WriteFile(favoritesFile, data, 0644); err != nil {
            return fmt.Errorf("failed to create favorites file: %v", err)
        }
    }

    // Validate existing file
    if data, err := os.ReadFile(favoritesFile); err == nil {
        var favorites Favorite
        if err := json.Unmarshal(data, &favorites); err != nil {
            // File exists but is invalid, recreate it
            favorites = Favorite{Photos: []Photo{}}
            data, _ := json.MarshalIndent(favorites, "", "    ")
            if err := os.WriteFile(favoritesFile, data, 0644); err != nil {
                return fmt.Errorf("failed to recreate invalid favorites file: %v", err)
            }
        }
    }

    return nil
}

// Load favorites from JSON file
func LoadFavorites() (Favorite, error) {
    mutex.RLock()
    defer mutex.RUnlock()

    data, err := os.ReadFile(favoritesFile)
    if err != nil {
        return Favorite{}, fmt.Errorf("failed to read favorites file: %v", err)
    }

    var favorites Favorite
    if err := json.Unmarshal(data, &favorites); err != nil {
        return Favorite{}, fmt.Errorf("failed to unmarshal favorites: %v", err)
    }

    return favorites, nil
}

// Save favorites to JSON file
func SaveFavorites(favorites Favorite) error {
    mutex.Lock()
    defer mutex.Unlock()

    data, err := json.MarshalIndent(favorites, "", "    ")
    if err != nil {
        return fmt.Errorf("failed to marshal favorites: %v", err)
    }

    if err := os.WriteFile(favoritesFile, data, 0644); err != nil {
        return fmt.Errorf("failed to save favorites: %v", err)
    }

    return nil
}

// Add photo to favorites
func AddToFavorites(photo Photo) error {
    favorites, err := LoadFavorites()
    if err != nil {
        return fmt.Errorf("failed to load favorites: %v", err)
    }

    // Check if photo already exists
    for _, p := range favorites.Photos {
        if p.ID == photo.ID {
            return nil // Photo already in favorites
        }
    }

    favorites.Photos = append(favorites.Photos, photo)
    return SaveFavorites(favorites)
}

// Remove photo from favorites
func RemoveFromFavorites(photoID int) error {
    favorites, err := LoadFavorites()
    if err != nil {
        return fmt.Errorf("failed to load favorites: %v", err)
    }

    // Filter out the photo to remove
    var newPhotos []Photo
    for _, p := range favorites.Photos {
        if p.ID != photoID {
            newPhotos = append(newPhotos, p)
        }
    }

    favorites.Photos = newPhotos
    return SaveFavorites(favorites)
}