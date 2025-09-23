package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type Data struct {
	NewContent PageContent
	Lang       string
	Form       subscribeFormValues  `json:"form"`
	Errors     []subscribeFormError `json:"errors"`
}

type subscribeFormValues struct {
	Subscribe string `json:"subscribe"`
}

type subscribeFormError struct {
	Field   string
	Message string
}

// type PageData struct {
// 	NewContent PageContent               `json:"NewContent"`
// 	Content    map[string]string         `json:"pageContent"`
// 	Lang       string                    `json:"lang"`
// 	Faq        []FrequentlyAskedQuestion `json:"faq"`
// 	Form       formValues                `json:"form"`
// 	Errors     []formError               `json:"errors"`
// }

func NotificationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		siteLang, _ := getSiteLanguage(r)
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

		form := subscribeFormValues{
			Subscribe: r.FormValue("subscribe"),
		}

		var formErrors []subscribeFormError

		if form.Subscribe == "" {
			var message string
			if siteLang == "zh" {
				message = "需要電子郵件。"
			} else {
				message = "Email is required."
			}
			formErrors = append(
				formErrors,
				subscribeFormError{
					Field:   "subscribe",
					Message: message,
				})
		}

		if form.Subscribe != "" {
			emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			matched, err := regexp.MatchString(emailRegex, form.Subscribe)
			if err != nil || !matched {
				var message string
				if siteLang == "zh" {
					message = "電子郵件地址無效"
				} else {
					message = "Invalid email address"
				}
				formErrors = append(
					formErrors,
					subscribeFormError{
						Field:   "subscribe",
						Message: message,
					})
			}
		}

		pageData := Data{
			NewContent: pageContentJSON,
			Lang:       siteLang,
			Form:       subscribeFormValues{},
			Errors:     formErrors,
		}

		err = templates.ExecuteTemplate(w, "subscribe", pageData)
		if err != nil {
			fmt.Println("Error rendering index template", err)
			http.Error(w, "Error rendering index template", http.StatusInternalServerError)
		}

	default:
		fmt.Println("Metrhod not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
