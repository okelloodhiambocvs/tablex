package service

import "tablex/backend/store"

// SimulatePayment confirms booking and locks table
func SimulatePayment(tableID int, bookingID int) bool {

	// lock table
	store.MarkTableUnavailable(tableID)

	// update booking status
	store.UpdateBookingStatus(bookingID, "CONFIRMED")

	return true
}