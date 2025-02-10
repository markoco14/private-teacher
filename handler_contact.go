package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type formValues struct {
	Name    string
	Email   string
	Message string
}

type formError struct {
	Field   string
	Message string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	
	lang := strings.TrimPrefix(r.URL.Path, "/")
	fmt.Println(lang)
	if lang != "en/contact" {
		lang = "zh"
	} else {
		lang = "en"
	}

	form := formValues{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	errors := validateFormData(form, lang)

	if len(errors) > 0 {
		type responseWithErrors struct {
			Form   formValues
			Errors []formError
			Lang string
		}


		data := responseWithErrors{
			Form:   form,
			Errors: errors,
			Lang:  lang,
		}

		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering form template", http.StatusInternalServerError)
		}

		return
	}

	err := emailTeacherMark(form)
	if err != nil {
		http.Error(w, "Something went wrong on our server. Please try again.", http.StatusInternalServerError)
		return
	}

	err = emailPotentialStudent(form, lang)
	if err != nil {
		// because email to Teacher successfully sent, don't need to break out
		// just log the error and continue
		log.Printf("Error sending email to potential student: %v", err)
	}

	if r.Header.Get("Hx-Request") == "true" {
		w.Header().Set("Hx-Trigger", `{"formSuccess": {"message": "Thank you for your message. We will get back to you soon."}}`)

		data := map[string]string{
			"Lang": lang,
		}
		err := templates.ExecuteTemplate(w, "form", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}

	if lang == "en" {
		http.Redirect(w, r, "/en", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func validateFormData(form formValues, lang string) []formError {
	var errors []formError

	if form.Name == "" {
		if lang == "en" {
			errors = append(errors, formError{Field: "name", Message: "Name is required"})
		} else {
			errors = append(errors, formError{Field: "name", Message: "名字是必需的"})
		}
	}

	if form.Email == "" {
		if lang == "en" {
			errors = append(errors, formError{Field: "email", Message: "Email is required"})
		} else {
			errors = append(errors, formError{Field: "email", Message: "電子郵件是必需的"})
		}
	}

	if form.Email != "" {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, err := regexp.MatchString(emailRegex, form.Email)
		if err != nil || !matched {
			if lang == "en" {
				errors = append(errors, formError{Field: "email", Message: "Invalid email address"})
			} else {
				errors = append(errors, formError{Field: "email", Message: "無效的電子郵件地址"})
			}
		}
	}

	if form.Message == "" {
		if lang == "en" {
			errors = append(errors, formError{Field: "message", Message: "Message is required"})
		} else {
			errors = append(errors, formError{Field: "message", Message: "訊息是必需的"})
		}
	} else if len(form.Message) < 10 {
		if lang == "en" {
			errors = append(errors, formError{Field: "message", Message: "Message must be at least 10 characters"})
		} else {
			errors = append(errors, formError{Field: "message", Message: "訊息必須至少10個字符"})
		}
	} else if len(form.Message) > 1000 {
		if lang	== "en" {
			errors = append(errors, formError{Field: "message", Message: "Message must be shorter than 1000 characters"})
		} else {
			errors = append(errors, formError{Field: "message", Message: "訊息必須少於1000個字符"})
		}
	}

	return errors
}

func getSMPTInfo() (string, int, string, string, error) {
	smtpEndpoint := os.Getenv("SMTP_ENDPOINT")
	if smtpEndpoint == "" {
		log.Fatalf("Something went wrong with our SMTP endpoint. Please try again.")
		return "", 0, "", "", fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Something went wrong with our SMTP port. Please try again.")
		return "", 0, "", "", fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	smtpUsername := os.Getenv("SMTP_USERNAME")
	if smtpUsername == "" {
		log.Fatalf("Something went wrong with our SMTP username. Please try again.")
		return "", 0, "", "", fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword == "" {
		log.Fatalf("Something went wrong with our SMTP password. Please try again.")
		return "", 0, "", "", fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	return smtpEndpoint, smtpPort, smtpUsername, smtpPassword, nil
}

func emailTeacherMark(form formValues) error {
	// Hardcoded email addresses
	infoEmail := os.Getenv("INFO_EMAIL")
	customerEmail := os.Getenv("TEST_CUSTOMER_EMAIL")
	// customerEmail := form.Email

	// load the email tempalte
	tmpl, err := template.ParseFiles("./templates/email-to-info.html.go")
	if err != nil {
		log.Fatalf("Error loading email template: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	// create a buffer to write the template to
	var body bytes.Buffer
	if err := tmpl.Execute(&body, form); err != nil {
		log.Fatalf("Error writing to email template: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	// send email
	message := gomail.NewMessage()
	message.SetHeader("From", infoEmail)         // email always comes from info@
	message.SetHeader("To", infoEmail)           // we are emailing ourself for our records and so we can reply back
	message.SetHeader("Reply-To", customerEmail) // the is the customer's email from the form
	message.SetHeader("Subject", "Teacher Mark Contact Form")
	message.SetBody("text/html", body.String())

	smtpEndpoint, smtpPort, smtpUsername, smtpPassword, err := getSMPTInfo()
	if err != nil {
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	d := gomail.NewDialer(smtpEndpoint, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(message); err != nil {
		log.Fatalf("Something went wrong with dial and send: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	return nil
}

func emailPotentialStudent(form formValues, lang string) error {
	infoEmail := os.Getenv("INFO_EMAIL")
	if infoEmail == "" {
		log.Fatalf("Something went wrong getting our info email.")
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	customerEmail := form.Email

	tmpl, err := template.ParseFiles("./templates/email-to-customer.html.go")
	if err != nil {
		log.Fatalf("Error loading email template: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}
	data := map[string]string{
		"Name": form.Name,
		"Lang": lang,
	}
	
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Fatalf("Error writing to email template: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	subject := "Thank you for contacting Teacher Mark"
	if lang == "zh" {
		subject = "感謝您聯繫Teacher Mark"
	}
	// prepare email info
	message := gomail.NewMessage()
	message.SetHeader("From", infoEmail) // email always comes from info@markoco14
	message.SetHeader("To", customerEmail)
	message.SetHeader("Reply-To", infoEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body.String())

	smtpEndpoint, smtpPort, smtpUsername, smtpPassword, err := getSMPTInfo()
	if err != nil {
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	d := gomail.NewDialer(smtpEndpoint, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(message); err != nil {
		log.Fatalf("Something went wrong with dial and send: %v", err)
		return fmt.Errorf("Something went wrong on our server. Please try again.")
	}

	return nil
}
