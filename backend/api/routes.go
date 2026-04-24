package api

import "net/http"

// RegisterRoutes sets up TABLEX endpoints
func RegisterRoutes() {
	http.HandleFunc("/clubs", GetClubsHandler)

	http.HandleFunc("/clubs/", func(w http.ResponseWriter, r *http.Request) {

		if len(r.URL.Path) > 0 && contains(r.URL.Path, "/tables") {
			GetClubTablesHandler(w, r)
			return
		}

		GetClubHandler(w, r)
	})
}

// helper function
func contains(path string, sub string) bool {
	return len(path) >= len(sub) && path[len(path)-len(sub):] == sub
}