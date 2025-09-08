package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type FrequentlyAskedQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type PageData struct {
	NewContent PageContent               `json:"NewContent"`
	Content    map[string]string         `json:"pageContent"`
	Lang       string                    `json:"lang"`
	Faq        []FrequentlyAskedQuestion `json:"faq"`
	Form       formValues                `json:"form"`
	Errors     []formError               `json:"errors"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var lang string
	cookie, err := r.Cookie("SITE_LANG")
	if errors.Is(err, http.ErrNoCookie) {
		lang = "en"
		newCookie := http.Cookie{
			Name:     "SITE_LANG",
			Value:    "en",
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   true,
		}
		http.SetCookie(w, &newCookie)
	} else {
		lang = cookie.Value
	}

	heroContentLocation := "./static/content/homepage.en.json"
	var pageContentJSON PageContent
	fileBytes, _ := os.ReadFile(heroContentLocation)
	err = json.Unmarshal(fileBytes, &pageContentJSON)
	if err != nil {
		fmt.Println("error getting json content")
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

	if lang == "en" {
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
		Lang:       lang,
		Faq:        faqList,
		Form:       formValues{},
		Errors:     []formError{},
	}

	err = templates.ExecuteTemplate(w, "index.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
