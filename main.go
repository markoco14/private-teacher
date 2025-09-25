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
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.HandleFunc("/classes", English)
	// http.HandleFunc("/coding", Coding)
	mux.HandleFunc("/notifications", NotificationsHandler)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
