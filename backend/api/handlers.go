package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"tablex/backend/service"
	"tablex/backend/core"
)

// Get all clubs
func GetClubsHandler(w http.ResponseWriter, r *http.Request) {
	data := service.GetAllClubs()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Get single club
func GetClubHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/clubs/")
	id, _ := strconv.Atoi(idStr)

	data := service.GetClubByID(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Get tables for a club
func GetClubTablesHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/clubs/")
	parts := strings.Split(idStr, "/")

	clubID, _ := strconv.Atoi(parts[0])

	club := service.GetClubByID(clubID)
	if club == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(club.(core.Club).Tables)
}