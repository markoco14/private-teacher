package main

import (
  "fmt"
  "net/http"
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
  fmt.Fprintf(w, "Maybe it is.")
}

