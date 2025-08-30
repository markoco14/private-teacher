package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Question struct {
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

type Answer struct {
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

type FrequentlyAskedQuestion struct {
	Question Question `json:"question"`
	Answer   Answer `json:"answer"`
}

type PageData struct {
	Content map[string]string		`json:"pageContent"`
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

func getPageContent(location string) (map[string]string, error) {
	file, err := os.Open(location)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err 
	}
	defer file.Close()
	
	contentMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	var currentHeader string
	for scanner.Scan() {
		line := strings.TrimSpace((scanner.Text()))
		if line == "" {
			continue // skip empty lines
		}

		// check if line has header
		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "---") {
			// extract the header key
			currentHeader = strings.Trim(line, "-")
			currentHeader = strings.ToLower(currentHeader)
		} else {
			// append the content to the current header's value
			contentMap[currentHeader] += line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading content from file: %v", err)
		return nil, err
	}

	return contentMap, nil
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

	faqContent, err := loadFrequentlyAsked()
	if err != nil {
		fmt.Println("Error loading FAQs", err)
		faqContent = []FrequentlyAskedQuestion{}
	}

	pageData := PageData{
		Content: pageContent,
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
