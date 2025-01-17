package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"gopkg.in/gomail.v2"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type formValues struct {
		Name    string
		Email   string
		Message string
	}

	type formError struct {
		Field   string
		Message string
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	form := formValues{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	var errors []formError

	if form.Name == "" {
		errors = append(errors, formError{Field: "name", Message: "Name is required"})
	}

	if form.Email == "" {
		errors = append(errors, formError{Field: "email", Message: "Email is required"})
	}

	if form.Email != "" {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, err := regexp.MatchString(emailRegex, form.Email)
		if err != nil || !matched {
			errors = append(errors, formError{Field: "email", Message: "Invalid email address"})
		}
	}

	if form.Message == "" {
		errors = append(errors, formError{Field: "message", Message: "Message is required"})
	} else if len(form.Message) < 10 {
		errors = append(errors, formError{Field: "message", Message: "Message must be at least 10 characters"})
	} else if len(form.Message) > 1000 {
		errors = append(errors, formError{Field: "message", Message: "Message must be shorter than 1000 characters"})
	}

	if len(errors) > 0 {
		type responseWithErrors struct {
			Form   formValues
			Errors []formError
		}

		data := responseWithErrors{
			Form:   form,
			Errors: errors,
		}

		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}

		return
	}

	// Hardcoded email addresses
	fromEmail := os.Getenv("FROM_EMAIL")
	toEmail := os.Getenv("TEST_EMAIL")

	// send email
	message := gomail.NewMessage()
	message.SetHeader("From", fromEmail)
	message.SetHeader("To", toEmail)
	message.SetHeader("Reply-To", fromEmail)
	message.SetHeader("Subject", "Teacher Mark Contact Form")
	message.SetBody("text/html", fmt.Sprintf("Name: %s<br>Email: %s<br>Message: %s", form.Name, form.Email, form.Message))

	smtpEndpoint := os.Getenv("SMTP_ENDPOINT")
	if smtpEndpoint == "" {
		http.Error(w, "Something went wrong with our email server. Please try again.", http.StatusInternalServerError)
		return
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		http.Error(w, "Something went wrong with our email server. Please try again.", http.StatusInternalServerError)
		return
	}

	smtpUsername := os.Getenv("SMTP_USERNAME")
	if smtpUsername == "" {
		http.Error(w, "Something went wrong with our email server. Please try again.", http.StatusInternalServerError)
		return
	}

	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword == "" {
		http.Error(w, "Something went wrong with our email server. Please try again.", http.StatusInternalServerError)
		return
	}

	d := gomail.NewDialer(smtpEndpoint, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(message); err != nil {
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Hx-Request") == "true" {
		w.Header().Set("Hx-Trigger", `{"formSuccess": {"message": "Thank you for your message. We will get back to you soon."}}`)
		err := templates.ExecuteTemplate(w, "form", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
