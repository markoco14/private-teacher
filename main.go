package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	// Routes
	http.HandleFunc("/", homeHandler)

	// http.HandleFunc()

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	type formValues struct {
		Name  string
		Email string
	}

	type formError struct {
		Field   string
		Message string
	}

	if r.Method == http.MethodGet {
		data := map[string]interface{}{
			"Form":   formValues{},
			"Errors": nil,
		}
		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		form := formValues{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
		}

		var errors []formError

		if form.Name == "" {
			errors = append(errors, formError{Field: "name", Message: "Name is required"})
		}

		if form.Email == "" {
			errors = append(errors, formError{Field: "email", Message: "Email is required"})
		}

		if len(errors) > 0 {
			data := map[string]interface{}{
				"Form":   form,
				"Errors": errors,
			}
			w.WriteHeader(http.StatusOK)
			if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}

    // if the form succeeds
    http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
