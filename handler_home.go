package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type FrequentlyAskedQuestion struct {
	Question string 	`json:"question"`
	Answer   string 	`json:"answer"`
}

type PageData struct {
	Content map[string]string		`json:"pageContent"`
	Lang string						`json:"lang"`
	Faq  []FrequentlyAskedQuestion	`json:"faq"`
	Form formValues					`json:"form"`
	Errors []formError				`json:"errors"`
}

func getFaqContent(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", location, err)
	}
	return string(content), nil
}

func createFaqList(content string) ([]FrequentlyAskedQuestion, error) {
	var faqs []FrequentlyAskedQuestion
	lines := strings.Split(content, "\n")

	var currentFAQ *FrequentlyAskedQuestion
	var isQuestion bool

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue // skip empty lines
		}

		if strings.HasPrefix(trimmedLine, "---") && strings.HasSuffix(trimmedLine, "---") {
			header := strings.Trim(trimmedLine, "-")
			if strings.HasPrefix(header, "Q_") {
				if currentFAQ != nil {
					faqs = append(faqs, *currentFAQ)
				}
				currentFAQ = &FrequentlyAskedQuestion{}
				isQuestion = true
			} else {
				isQuestion = false
			}
		} else {
			if currentFAQ != nil {
				if isQuestion {
					currentFAQ.Question += trimmedLine
				} else {
					if currentFAQ.Answer != "" {
						currentFAQ.Answer += " "
					}
					currentFAQ.Answer += trimmedLine
				}
			}
		}
	}

	if currentFAQ != nil {
		faqs = append(faqs, *currentFAQ)
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
		Content: pageContent,
		Lang: lang,
		Faq: faqList,
		Form: formValues{},
		Errors: []formError{},
	}
	
	err = templates.ExecuteTemplate(w, "index.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
