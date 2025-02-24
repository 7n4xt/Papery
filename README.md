# Papery

![Papery Logo](/assets/Icons/paperyLogo.png)

A modern web application for discovering, managing, and enjoying high-quality wallpapers.

## 🌟 Overview

Papery helps users discover and organize beautiful wallpapers through an intuitive interface. Browse thousands of high-resolution images, create personalized collections, and transform your digital spaces with stunning visuals.

## ✨ Key Features

### Collections & Favorites
- Create custom collections to organize wallpapers by theme, mood, or occasion
- Quick favorite/unfavorite functionality with one-click actions
- Smart recommendations based on your favorite styles

### Premium Content Library
- Access to thousands of high-resolution wallpapers
- Regular updates with fresh content
- Support for multiple devices (desktop, tablet, TV's lol)
- Multiple resolution options for each wallpaper

### Powerful Search Engine
- Advanced filtering by resolution, orientation, color palette
- Instant search
- Custom sorting options (orientation, size, most color)

## 🛠️ Technology Stack

### Frontend
- **HTML5** - Semantic structure and accessibility
- **CSS3** - Custom animations and responsive layouts
- **JavaScript** - Interactive elements and dynamic content loading

### Backend
- **GoLang** - High-performance API server and image processing
- **RESTful API** - Clean architecture for seamless frontend-backend communication

### Data Storage
- **JSON** - Efficient metadata storage and retrieval

## 📋 Prerequisites

- Git ([Download](https://git-scm.com/downloads))
- Go 1.23.4 or higher ([Download](https://go.dev/dl/))
- Web browser with JavaScript enabled

## 🚀 Getting Started

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/YourUsername/Papery.git
   cd Papery
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Build and run the application
   ```bash
   go build
   ./papery
   ```

4. Access in browser
   ```
   http://localhost:8080
   ```

## 🧩 Project Structure

```
papery/
├── main.go                      # Application entry point
├── go.mod                       # Go module file
├── README.md                    # Project documentation
│
├── assets/                      # Static assets
│   ├── CSS/
│   │   ├── universal.css        # Universal styles
│   │   ├── detailsPage.css      # Details page styles
│   │   ├── home.css             # Home page styles
│   │   ├── search.css           # Search page styles
│   │   └── favorite.css         # Favorites page styles
│   ├── Fonts/
│   │   └── BrightMarching-Regular.otf  # Custom font
│   ├── Icons/                   # Application icons
│   └── js/
│       └── favorites.js         # Favorites functionality
│
├── controllers/                 # MVC Controllers
│   ├── about.controller.go      # About page controller
│   ├── details.controller.go    # Details page controller
│   ├── favorite.controller.go   # Favorites page controller
│   ├── home.controller.go       # Home page controller
│   └── search.controller.go     # Search page controller
│
├── data/                        # Data storage
│   └── favorites.json           # User favorites data
│
├── routes/                      # Routing configuration
│   ├── about.route.go           # About page routes
│   ├── details.route.go         # Details page routes
│   ├── favorite.route.go        # Favorites page routes
│   ├── home.route.go            # Home page routes 
│   ├── main.routes.go           # Main routing setup
│   └── search.route.go          # Search page routes
│
├── services/                    # Business logic
│   ├── details.service.go       # Photo details service
│   ├── favorites.service.go     # Favorites management
│   ├── home.services.go         # Home page content
│   └── search.service.go        # Search functionality
│
└── templates/                   # HTML templates
    ├── about.html               # About page template
    ├── details.html             # Details page template
    ├── favorite.html            # Favorites page template
    ├── home.html                # Home page template
    ├── search.html              # Search page template
    └── templates.go             # Template handler
```

## 🤝 Author

ESUGHI Abdulmalek

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- All the amazing photographers who provide beautiful wallpapers
- The open-source community for their invaluable tools and libraries, 

# Merci beaucoup!