package service

import (
	"tablex/backend/core"
	"tablex/backend/store"
)

// BookTable handles booking logic with split suggestion
func BookTable(user string, table core.Table, people int, budget int) core.Booking {

	status := "CONFIRMED"

	// if budget is low, mark pending (possible split)
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