package main

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	err := templates.ExecuteTemplate(w, "index.html.go", nil)
	if err != nil {
		http.Error(w, "Error rendering index template", http.StatusInternalServerError)
	}
}
