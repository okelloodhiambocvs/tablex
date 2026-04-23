package api

import (
	"encoding/json"
	"net/http"

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

// SplitRequestHandler handles split bill creation
func SplitRequestHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		TableID int
		UserA   string
		UserB   string
		Price   int
	}

	// decode incoming request
	json.NewDecoder(r.Body).Decode(&req)

	// call service layer
	split := service.CreateSplitSuggestion(
		req.TableID,
		req.UserA,
		req.UserB,
		req.Price,
	)

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(split)
}

// PaymentHandler simulates payment and locks the table
func PaymentHandler(w http.ResponseWriter, r *http.Request) {

	// only allow POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// request body structure
	var req struct {
	TableID   int
	BookingID int
}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate table exists
	table, found := store.GetTableByID(req.TableID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "table not found",
		})
		return
	}

	// check availability
	if !table.Available {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "table already booked",
		})
		return
	}

	// simulate payment and lock table
	success := service.SimulatePayment(req.TableID, req.BookingID)

	response := map[string]interface{}{
		"success": success,
		"message": "payment successful, table booked",
		"tableID": req.TableID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}