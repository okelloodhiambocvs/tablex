package main

import (
	"fmt"
	"net/http"
	"tablex/backend/api"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TABLEX API is running")
}

func main() {
	api.RegisterRoutes()

	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}