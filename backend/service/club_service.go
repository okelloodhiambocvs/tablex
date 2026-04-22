package service

import "tablex/backend/store"

// GetAllClubs returns all clubs
func GetAllClubs() interface{} {
	return store.GetClubs()
}

// GetClubByID returns a single club
func GetClubByID(id int) interface{} {
	clubs := store.GetClubs()

	for _, c := range clubs {
		if c.ID == id {
			return c
		}
	}
	return nil
}