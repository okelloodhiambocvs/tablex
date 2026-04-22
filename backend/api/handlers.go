package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tablex/backend/core"
	"tablex/backend/service"
	"tablex/backend/store"
)

// BookTableHandler handles booking requests
func BookTableHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		UserName string
		TableID  int
		People   int
		Budget   int
	}

	json.NewDecoder(r.Body).Decode(&req)

	// find table
	clubs := store.GetClubs()

	var selectedTable core.Table

	for _, c := range clubs {
		for _, t := range c.Tables {
			if t.ID == req.TableID {
				selectedTable = t
			}
		}
	}

	booking := service.BookTable(req.UserName, selectedTable, req.People, req.Budget)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}