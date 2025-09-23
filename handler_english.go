package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type kidsContent struct {
	Headline string
	First    string
	Second   string
	Third    string
	Speak    template.HTML
	Read     template.HTML
	Write    template.HTML
	Listen   template.HTML
}

type adultsContent struct {
	Headline    string
	Description string
	SubheadOne  string `json:"subhead_one"`
	SupportOne  string `json:"support_one"`
	SubheadTwo  string `json:"subhead_two"`
	SupportTwo  string `json:"support_two"`
}

type businessContent struct {
	Headline     string
	Description  string
	SupportOne   string `json:"support_one"`
	SupportTwo   string `json:"support_two"`
	SupportThree string `json:"support_three"`
	SupportFour  string `json:"support_four"`
}

type testContent struct {
	Headline         string
	Ielts            string
	IeltsDescription string `json:"ielts_description"`
	Gept             string
	GeptDescription  string `json:"gept_description"`
	Other            string
	OtherDescription string `json:"other_description"`
}

type topContent struct {
	Headline     string
	Description  string
	KidsLink     string `json:"kids_link"`
	AdultsLink   string `json:"adults_link"`
	BusinessLink string `json:"business_link"`
	TestLink     string `json:"test_link"`
}

type pageContent struct {
	Base     BaseContent
	Top      topContent
	Kids     kidsContent
	Adults   adultsContent
	Business businessContent
	Test     testContent
}

type pageData struct {
	NewContent pageContent `json:"NewContent"`
	Lang       string      `json:"lang"`
}

func English(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/english" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	siteLang, cookieFound := getSiteLanguage(r)
	if !cookieFound {
		setSiteLanguageCookie(w, siteLang)
	}

	var heroContentLocation string
	if siteLang == "zh" {
		heroContentLocation = "./static/locales/zh/classes.json"
	} else {
		heroContentLocation = "./static/locales/en/classes.json"
	}

	var pageContentJSON pageContent
	fileBytes, _ := os.ReadFile(heroContentLocation)
	err := json.Unmarshal(fileBytes, &pageContentJSON)
	if err != nil {
		fmt.Println("error getting json content")
	}

	pageData := pageData{
		NewContent: pageContentJSON,
		Lang:       siteLang,
	}

	err = templates.ExecuteTemplate(w, "english.gohtml", pageData)
	if err != nil {
		fmt.Println("Error rendering index template", err)
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
