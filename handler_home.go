package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type FrequentlyAskedQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type PageData struct {
	Lang string						`json:"lang"`
	Faq  []FrequentlyAskedQuestion	`json:"faq"`
	Form formValues					`json:"form"`
	Errors []formError				`json:"errors"`
}

func loadFrequentlyAsked() ([]FrequentlyAskedQuestion, error) {
	data, err := os.ReadFile("./data/faq.json")
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}

	var faqs []FrequentlyAskedQuestion
	err = json.Unmarshal(data, &faqs)
	if err != nil {
		fmt.Println("Error unmarshalling JSON")
		return nil, err
	}

	return faqs, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lang := strings.TrimPrefix(r.URL.Path, "/")
	if lang != "en" {
		lang = "zh"
	}

	faqContent, err := loadFrequentlyAsked()
	if err != nil {
		fmt.Println("Error loading FAQs", err)
		faqContent = []FrequentlyAskedQuestion{}
	}

	pageData := PageData{
		Lang: lang,
		Faq: faqContent,
		Form: formValues{},
		Errors: []formError{},
	}
	

	

	err = templates.ExecuteTemplate(w, "index.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
