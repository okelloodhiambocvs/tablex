package main

import (
	"fmt"
	"net/http"
)

// simple health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TABLEX API is running")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	// start server
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}