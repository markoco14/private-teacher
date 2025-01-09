package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main() {
  // Routes
  http.HandleFunc("/", homeHandler)

  // Start the server
  fmt.Println("Server is running on port 8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  template := template.Must(template.ParseFiles("./templates/index.html"))
  if err := template.Execute(w, nil);
  err != nil {
    http.Error(w, "Error rendering template", http.StatusInternalServerError)
  }
}

