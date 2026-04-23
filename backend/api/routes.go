package api

import (
	"net/http"
	"strings"
)

// RegisterRoutes sets up all API endpoints
func RegisterRoutes() {

	// health check route
	http.HandleFunc("/health", HealthHandler)

	// get all clubs
	http.HandleFunc("/clubs", GetClubsHandler)

	// dynamic club routes
	http.HandleFunc("/clubs/", func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		// check if request is for tables
		if strings.Contains(path, "/tables") {
			GetClubTablesHandler(w, r)
			return
		}

		// otherwise return single club
		GetClubHandler(w, r)
	})

	// booking route
	http.HandleFunc("/book", BookTableHandler)

	// split bill route
	http.HandleFunc("/split", SplitRequestHandler)
}