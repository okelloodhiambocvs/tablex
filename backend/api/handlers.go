package api

import (
	"encoding/json"
	"net/http"
	"tablex/backend/store"
)

// GetClubsHandler returns all clubs
func GetClubsHandler(w http.ResponseWriter, r *http.Request) {
	clubs := store.GetClubs()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clubs)
}