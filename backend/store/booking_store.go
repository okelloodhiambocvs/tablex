package store

import "tablex/backend/core"

var bookings []core.Booking
var bookingID = 1

// SaveBooking stores a booking
func SaveBooking(b core.Booking) core.Booking {
	b.ID = bookingID
	bookingID++

	bookings = append(bookings, b)
	return b
}

// GetBookings returns all bookings
func GetBookings() []core.Booking {
	return bookings
}