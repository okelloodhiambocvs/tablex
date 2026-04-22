package api

import "net/http"

// RegisterRoutes sets up all endpoints
func RegisterRoutes() {
	http.HandleFunc("/clubs", GetClubsHandler)
}