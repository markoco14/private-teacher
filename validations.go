package main

import (
	"regexp"
)

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
		if lang == "en" {
			errors = append(errors, formError{Field: "message", Message: "Message must be shorter than 1000 characters"})
		} else {
			errors = append(errors, formError{Field: "message", Message: "訊息必須少於1000個字符"})
		}
	}

	return errors
}