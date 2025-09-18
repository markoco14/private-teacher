package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func English(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/english" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	siteLang, cookieFound := getSiteLanguage(r)
	if !cookieFound {
		newCookie := http.Cookie{
			Name:     "SITE_LANG",
			Value:    siteLang,
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   true,
		}
		http.SetCookie(w, &newCookie)
	}

	heroContentLocation := "./static/content/homepage.en.json"
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

	if siteLang == "en" {
		fileLocation = "./static/content/faq.en.txt"
	} else {
		fileLocation = "./static/content/faq.zh.txt"
	}

	faqContent, err := getFaqContent(fileLocation)
	if err != nil {
		log.Printf("Error getting faq content from file: %v", err)
		// Set the 500 status code
		w.WriteHeader(http.StatusInternalServerError)
		// Execute the error template
		templates.ExecuteTemplate(w, "error.gohtml", nil)
		return
	}

	faqList, err := createFaqList(faqContent)
	if err != nil {
		fmt.Println("Error loading FAQs", err)
	}

	pageData := PageData{
		NewContent: pageContentJSON,
		Content:    pageContent,
		Lang:       siteLang,
		Faq:        faqList,
		Form:       formValues{},
		Errors:     []formError{},
	}

	err = templates.ExecuteTemplate(w, "english.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
