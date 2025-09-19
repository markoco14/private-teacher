package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type Data struct {
	Form   subscribeFormValues  `json:"form"`
	Errors []subscribeFormError `json:"errors"`
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
		fmt.Println("post to notifications")

		form := subscribeFormValues{
			Subscribe: r.FormValue("subscribe"),
		}

		var formErrors []subscribeFormError

		if form.Subscribe == "" {
			formErrors = append(
				formErrors,
				subscribeFormError{
					Field:   "subscribe",
					Message: "Email is required.",
				})
		}

		if form.Subscribe != "" {
			emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
			matched, err := regexp.MatchString(emailRegex, form.Subscribe)
			if err != nil || !matched {
				formErrors = append(
					formErrors,
					subscribeFormError{
						Field:   "subscribe",
						Message: "Invalid email address",
					})
			}
		}

		fmt.Println(formErrors)

		pageData := Data{
			Form: subscribeFormValues{},
			Errors: formErrors,
		}

		err := templates.ExecuteTemplate(w, "subscribe", pageData)
		if err != nil {
			fmt.Println("Error rendering index template", err)
			http.Error(w, "Error rendering index template", http.StatusInternalServerError)
		}

	default:
		fmt.Println("Metrhod not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
