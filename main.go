package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func loadTemplates() (*template.Template, error) {
	tmpl := template.New("")
	tmpl, err := tmpl.ParseGlob("./templates/*.html.go")
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	_, err = tmpl.ParseGlob("./templates/partials/*.html.go")
	if err != nil {
		log.Fatalf("Error loading partials: %v", err)
	}

	return tmpl, nil
}

var templates = template.Must(loadTemplates())
// var templates = template.Must(template.ParseGlob("./templates/*.html.go"))

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// handle server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle favicon.ico requests explicitly
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
