package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FrequentlyAskedQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type heroContent struct {
	Headline   string `json:"headline"`
	Subhead    string `json:"subhead"`
	CtaClasses string `json:"cta_classes"`
	CtaTeacher string `json:"cta_teacher"`
}

type aboutContent struct {
	Headline string
	One      string
	Two      string
	Three    string
	Four     string
	Five     string
	Six      string
	Seven    string
}

type classesContent struct {
	Headline            string
	Description         string
	Kids                string
	KidsDescription     string `json:"kids_description"`
	Adults              string
	AdultsDescription   string `json:"adults_description"`
	Business            string
	BusinessDescription string `json:"business_description"`
	Test                string
	TestDescription     string `json:"test_description"`
}

type comingSoonContent struct {
	Headline    string
	One         string
	Two         string
	Placeholder string
	Cta         string
	CtaPending  string `json:"cta_pending"`
	Disclaimer  string
}

type PageContent struct {
	Base    BaseContent
	Hero    heroContent
	Classes classesContent
	Coming  comingSoonContent
	About   aboutContent
	Contact ContactContent
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
		siteLang, cookieFound := getSiteLanguage(r)
		if !cookieFound {
			setSiteLanguageCookie(w, siteLang)
		}

		pageData := struct{ Lang string }{Lang: siteLang}

        w.WriteHeader(http.StatusNotFound)
		err := templates.ExecuteTemplate(w, "404.gohtml", pageData)
		if err != nil {
			fmt.Println("Error rendering index template", err)
			http.Error(w, "Error rendering index template", http.StatusInternalServerError)
		}
		
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	err = templates.ExecuteTemplate(w, "index.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
