package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getSiteLanguage( r *http.Request) (string, bool) {
	cookie, err := r.Cookie("SITE_LANG")
	if errors.Is(err, http.ErrNoCookie) {
		return "en", false
	}
	return cookie.Value, true
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