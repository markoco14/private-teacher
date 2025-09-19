package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadTemplates() (*template.Template, error) {
	tmpl := template.New("")
	tmpl, err := tmpl.ParseGlob("./templates/*.gohtml")
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	_, err = tmpl.ParseGlob("./templates/partials/*.gohtml")
	if err != nil {
		log.Fatalf("Error loading partials folder: %v", err)
	}

	_, err = tmpl.ParseGlob("./templates/ui/*.gohtml")
	if err != nil {
		log.Fatalf("Error loading ui folder: %v", err)
	}

	return tmpl, nil
}

var templates = template.Must(loadTemplates())

func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	// handle server static files
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})

	// Handle favicon.ico requests explicitly
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/english", English)
	http.HandleFunc("/coding", Coding)
	http.HandleFunc("/notifications", NotificationsHandler)


	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
