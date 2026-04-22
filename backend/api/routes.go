package api

import "net/http"

// RegisterRoutes initializes all TABLEX API endpoints
func RegisterRoutes() {

	// health check
	http.HandleFunc("/health", healthHandler)

	// clubs list
	http.HandleFunc("/clubs", GetClubsHandler)

	// booking endpoint
	http.HandleFunc("/book", BookTableHandler)

	// dynamic club routes (IMPORTANT)
	http.HandleFunc("/clubs/", clubRouter)
}

// clubRouter handles all /clubs/* routes
func clubRouter(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	// /clubs/{id}/tables
	if len(path) > len("/clubs/") && contains(path, "/tables") {
		GetClubTablesHandler(w, r)
		return
	}

	// /clubs/{id}
	GetClubHandler(w, r)
}

// simple helper to check substring
func contains(path string, sub string) bool {
	return len(path) >= len(sub) && (path[len(path)-len(sub):] == sub)
}