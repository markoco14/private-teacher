package main

import (
	"log"
	"net/http"
	"strings"
)

type formValues struct {
	Name    string
	Email   string
	Message string
}

type formError struct {
	Field   string
	Message string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	var lang string
	if strings.Contains(r.Header.Get("Referer"), "en") {
		lang = "en"
		} else {
		lang = "zh"
	}

	var fileLocation string
	if lang == "en" {
		fileLocation = "./static/content/homepage.en.txt"
	} else {
		fileLocation = "./static/content/homepage.zh.txt"
	}

	pageContent, err := getPageContent(fileLocation)
	if err != nil {
		log.Printf("Error getting page content from file: %v", err)
		// Set the 500 status code
        w.WriteHeader(http.StatusInternalServerError)
        // Execute the error template
        templates.ExecuteTemplate(w, "error.gohtml", nil)
        return
	}

	form := formValues{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	errors := validateFormData(form, lang)

	if len(errors) > 0 {
		type responseWithErrors struct {
			Content map[string]string
			Form   formValues
			Errors []formError
			Lang   string
		}

		data := responseWithErrors{
			Content: pageContent,
			Form:   form,
			Errors: errors,
			Lang:   lang,
		}

		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering form template", http.StatusInternalServerError)
		}

		return
	}

	err = emailTeacherMark(form)
	if err != nil {
		http.Error(w, "Something went wrong on our server. Please try again.", http.StatusInternalServerError)
		return
	}

	err = emailPotentialStudent(form, lang)
	if err != nil {
		// because email to Teacher successfully sent, don't need to break out
		// just log the error and continue
		log.Printf("Error sending email to potential student: %v", err)
	}

	if r.Header.Get("Hx-Request") == "true" {
		w.Header().Set("Hx-Trigger", `{"formSuccess": {"message": "Thank you for your message. We will get back to you soon."}}`)

		data := map[string]string{
			"Lang": lang,
		}
		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}

	if lang == "en" {
		http.Redirect(w, r, "/en", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
