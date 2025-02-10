package main

import (
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lang := strings.TrimPrefix(r.URL.Path, "/")
	if lang != "en" {
		lang = "zh"
	}

	w.WriteHeader(http.StatusOK)

	data := map[string]string{
		"Lang": lang,
	}
	
	err := templates.ExecuteTemplate(w, "index.html.go", data)
	if err != nil {
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
