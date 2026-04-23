package store

import "tablex/backend/core"

// global in-memory clubs storage
var clubs = []core.Club{
	{
		ID:       1,
		Name:     "Berlin Lounge",
		Location: "Kisumu Mamboleo",
		Tables: []core.Table{
			{ID: 1, ClubID: 1, Type: "VVIP", Price: 15000, Available: true},
			{ID: 2, ClubID: 1, Type: "VIP", Price: 8000, Available: true},
			{ID: 3, ClubID: 1, Type: "REGULAR", Price: 3000, Available: true},
		},
	},
	{
		ID:       2,
		Name:     "VIP Lounge",
		Location: "Milimani",
		Tables: []core.Table{
			{ID: 4, ClubID: 2, Type: "VVIP", Price: 20000, Available: true},
			{ID: 5, ClubID: 2, Type: "VIP", Price: 10000, Available: true},
			{ID: 6, ClubID: 2, Type: "REGULAR", Price: 4000, Available: true},
		},
	},
}

// GetClubs returns all clubs
func GetClubs() []core.Club {
	return clubs
}

// GetTableByID finds a table by ID
func GetTableByID(tableID int) (*core.Table, bool) {
	for i := range clubs {
		for j := range clubs[i].Tables {

			if clubs[i].Tables[j].ID == tableID {
				return &clubs[i].Tables[j], true
			}
		}
	}
	return nil, false
}

// MarkTableUnavailable sets table availability to false
func MarkTableUnavailable(tableID int) bool {

	for i := range clubs {
		for j := range clubs[i].Tables {

			if clubs[i].Tables[j].ID == tableID {

				// check if already booked
				if !clubs[i].Tables[j].Available {
					return false
				}

				clubs[i].Tables[j].Available = false
				return true
			}
		}
	}
	return false
}