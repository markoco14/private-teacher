package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	// handle server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle favicon.ico requests explicitly
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type formValues struct {
		Name    string
		Email   string
		Message string
	}

	type formError struct {
		Field   string
		Message string
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	form := formValues{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	var errors []formError

	if form.Name == "" {
		errors = append(errors, formError{Field: "name", Message: "Name is required"})
	}

	if form.Email == "" {
		errors = append(errors, formError{Field: "email", Message: "Email is required"})
	}

	if form.Email != "" {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, err := regexp.MatchString(emailRegex, form.Email)
		if err != nil || !matched {
			errors = append(errors, formError{Field: "email", Message: "Invalid email address"})
		}
	}

	if form.Message == "" {
		errors = append(errors, formError{Field: "message", Message: "Message is required"})
	} else if len(form.Message) < 10 {
		errors = append(errors, formError{Field: "message", Message: "Message must be at least 10 characters"})
	} else if len(form.Message) > 1000 {
		errors = append(errors, formError{Field: "message", Message: "Message must be shorter than 1000 characters"})
	}
	
	
	if len(errors) > 0 {
		type responseWithErrors struct {
			Form   formValues
			Errors []formError
		}

		data := responseWithErrors{
			Form:   form,
			Errors: errors,
		}
		
		w.WriteHeader(http.StatusOK)
		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
		
		return
	}
	
	w.Header().Set("Hx-Trigger", "formSuccess")
	w.WriteHeader(http.StatusOK)



	err := templates.ExecuteTemplate(w, "form", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}
