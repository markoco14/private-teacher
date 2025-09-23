package main

type BaseContent struct {
	Logo   string
	Rights string
}

type ContactContent struct {
	Headline    string
	Description string
	Line        string
	Email       string
	FormName    string `json:"form_name"`
	FormEmail   string `json:"form_email"`
	FormMessage string `json:"form_message"`
	Cta         string
	CtaPending  string `json:"cta_pending"`
}
