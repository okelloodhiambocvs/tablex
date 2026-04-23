package api

import "net/http"

// RegisterRoutes sets up all TABLEX endpoints
func RegisterRoutes() {

	// get all clubs
	http.HandleFunc("/clubs", GetClubsHandler)

	// dynamic club routes
	http.HandleFunc("/clubs/", func(w http.ResponseWriter, r *http.Request) {

		// check if request is for tables
		if contains(r.URL.Path, "/tables") {
			GetClubTablesHandler(w, r)
			return
		}

		// otherwise return single club
		GetClubHandler(w, r)
	})

	// booking endpoint
	http.HandleFunc("/book", BookTableHandler)

	// split bill endpoint
	http.HandleFunc("/split", SplitRequestHandler)

	// payment endpoint (THIS IS STEP 6.4)
	http.HandleFunc("/pay", PaymentHandler)
}

// simple helper function
func contains(path string, sub string) bool {
	return len(path) >= len(sub) && path[len(path)-len(sub):] == sub
}