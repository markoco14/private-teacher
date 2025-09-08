package main


type BaseContent struct {
	Logo string
	Rights string
}

type HeroContent struct {
	Headline   string `json:"headline"`
	Subhead    string `json:"subhead"`
	CtaClasses string `json:"cta_classes"`
	CtaTeacher string `json:"cta_teacher"`
}

type PageContent struct {
	Base BaseContent
	Hero HeroContent
}