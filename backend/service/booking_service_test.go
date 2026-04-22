package service

import (
	"testing"
	"tablex/backend/core"
)

func TestBookTable(t *testing.T) {

	table := core.Table{
		ID:    1,
		Type:  "VIP",
		Price: 5000,
	}

	booking := BookTable("John", table, 4, 3000)

	if booking.Status != "PENDING" {
		t.Error("Expected PENDING when budget is low")
	}
}