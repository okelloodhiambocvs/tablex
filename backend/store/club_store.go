package store

import "tablex/backend/core"

// GetClubs returns sample clubs in Kisumu
func GetClubs() []core.Club {
	return []core.Club{
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
}