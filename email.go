package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func getSMPTInfo() (string, int, string, string, error) {
	smtpEndpoint := os.Getenv("SMTP_ENDPOINT")
	if smtpEndpoint == "" {
		log.Printf("Something went wrong with our SMTP endpoint. Please try again.")
		return "", 0, "", "", fmt.Errorf("something went wrong on our server, please try again")
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Printf("Something went wrong with our SMTP port. Please try again.")
		return "", 0, "", "", fmt.Errorf("something went wrong on our server, please try again")
	}

	smtpUsername := os.Getenv("SMTP_USERNAME")
	if smtpUsername == "" {
		log.Printf("Something went wrong with our SMTP username. Please try again.")
		return "", 0, "", "", fmt.Errorf("something went wrong on our server, please try again")
	}

	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword == "" {
		log.Printf("Something went wrong with our SMTP password. Please try again.")
		return "", 0, "", "", fmt.Errorf("something went wrong on our server, please try again")
	}

	return smtpEndpoint, smtpPort, smtpUsername, smtpPassword, nil
}

func emailTeacherMark(form formValues) error {
	// Hardcoded email addresses
	infoEmail := os.Getenv("INFO_EMAIL")
	// customerEmail := os.Getenv("TEST_CUSTOMER_EMAIL")
	customerEmail := form.Email

	// load the email tempalte
	tmpl, err := template.ParseFiles("./templates/email-to-info.gohtml")
	if err != nil {
		log.Printf("Error loading email template: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	// create a buffer to write the template to
	var body bytes.Buffer
	if err := tmpl.Execute(&body, form); err != nil {
		log.Printf("Error writing to email template: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
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
		log.Printf("Something went wrong with dial and send: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	return nil
}

func emailPotentialStudent(form formValues, lang string) error {
	infoEmail := os.Getenv("INFO_EMAIL")
	if infoEmail == "" {
		log.Printf("Something went wrong getting our info email.")
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	customerEmail := form.Email

	tmpl, err := template.ParseFiles("./templates/email-to-customer.gohtml")
	if err != nil {
		log.Printf("Error loading email template: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
	}
	data := map[string]string{
		"Name": form.Name,
		"Lang": lang,
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Printf("Error writing to email template: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	subject := "Thank you for contacting Teacher Mark"
	if lang == "zh" {
		subject = "感謝您聯繫Teacher Mark"
	}
	// prepare email info
	message := gomail.NewMessage()
	message.SetHeader("From", infoEmail) // email always comes from info@teachermark.com.tw
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
		log.Printf("Something went wrong with dial and send: %v", err)
		return fmt.Errorf("something went wrong on our server, please try again")
	}

	return nil
}
