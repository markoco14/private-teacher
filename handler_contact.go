package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	switch r.Method {

	case http.MethodGet:
		siteLang, cookieFound := getSiteLanguage(r)
		if !cookieFound {
			setSiteLanguageCookie(w, siteLang)
		}

		var heroContentLocation string
		if siteLang == "zh" {
			heroContentLocation = "./static/locales/zh/home.json"
		} else {
			heroContentLocation = "./static/locales/en/home.json"
		}
		
		var pageContentJSON PageContent
		fileBytes, _ := os.ReadFile(heroContentLocation)
		err := json.Unmarshal(fileBytes, &pageContentJSON)
		if err != nil {
			fmt.Println("error getting json content")
		}

		var fileLocation string
		if siteLang == "en" {
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

		pageData := PageData{
			NewContent: pageContentJSON,
			Content:    pageContent,
			Lang:       siteLang,
			Form:       formValues{},
			Errors:     []formError{},
		}

		err = templates.ExecuteTemplate(w, "contact.gohtml", pageData)
		if err != nil {
			http.Error(w, "Error rendering form template", http.StatusInternalServerError)
		}

		return

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		siteLang, cookieFound := getSiteLanguage(r)
		if !cookieFound {
			setSiteLanguageCookie(w, siteLang)
		}

		var heroContentLocation string
		if siteLang == "zh" {
			heroContentLocation = "./static/locales/zh/home.json"
		} else {
			heroContentLocation = "./static/locales/en/home.json"
		}
		
		var pageContentJSON PageContent
		fileBytes, _ := os.ReadFile(heroContentLocation)
		err := json.Unmarshal(fileBytes, &pageContentJSON)
		if err != nil {
			fmt.Println("error getting json content")
		}

		var fileLocation string
		if siteLang == "en" {
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

		errors := validateFormData(form, siteLang)

		if len(errors) > 0 {
			type responseWithErrors struct {
				NewContent PageContent `json:"NewContent"`
				Content    map[string]string
				Form       formValues
				Errors     []formError
				Lang       string
			}

			data := responseWithErrors{
				NewContent: pageContentJSON,
				Content:    pageContent,
				Form:       form,
				Errors:     errors,
				Lang:       siteLang,
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

		err = emailPotentialStudent(form, siteLang)
		if err != nil {
			// because email to Teacher successfully sent, don't need to break out
			// just log the error and continue
			log.Printf("Error sending email to potential student: %v", err)
		}

		if r.Header.Get("Hx-Request") == "true" {
			w.Header().Set("Hx-Trigger", `{"formSuccess": {"message": "Thank you for your message. We will get back to you soon."}}`)

			type response struct {
				NewContent PageContent `json:"NewContent"`
				Content    map[string]string
				Form       formValues
				Errors     []formError
				Lang       string
			}

			data := response{
				NewContent: pageContentJSON,
				Content:    pageContent,
				Form:       formValues{},
				Errors:     errors,
				Lang:       siteLang,
			}

			err := templates.ExecuteTemplate(w, "form", data)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
		}

		if siteLang == "en" {
			http.Redirect(w, r, "/en", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
