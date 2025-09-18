package main

type BaseContent struct {
	Logo   string
	Rights string
}

type HeroContent struct {
	Headline   string `json:"headline"`
	Subhead    string `json:"subhead"`
	CtaClasses string `json:"cta_classes"`
	CtaTeacher string `json:"cta_teacher"`
}

type AboutContent struct {
	Headline string
	One      string
	Two      string
	Three    string
	Four     string
	Five     string
	Six      string
	Seven    string
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
type PageContent struct {
	Base  BaseContent
	Hero  HeroContent
	About AboutContent
	Contact ContactContent
}
