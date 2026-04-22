package service

import (
	"tablex/backend/core"
	"tablex/backend/store"
)

// BookTable handles booking logic
func BookTable(user string, table core.Table, people int, budget int) core.Booking {

	status := "CONFIRMED"

	// budget check
	if budget < table.Price {
		status = "PENDING"
	}

	booking := core.Booking{
		UserName: user,
		TableID:  table.ID,
		People:   people,
		Budget:   budget,
		Status:   status,
	}

	return store.SaveBooking(booking)
}