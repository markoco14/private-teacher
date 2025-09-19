package main

import (
	"fmt"
	"net/http"
	"time"
)

type Data struct {
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
		time.Sleep(5 * time.Second)

		pageData := Data{}

		err := templates.ExecuteTemplate(w, "subscribe.gohtml", pageData)
		if err != nil {
			fmt.Println("Error rendering index template", err)
			http.Error(w, "Error rendering index template", http.StatusInternalServerError)
		}

	default:
		fmt.Println("Metrhod not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
