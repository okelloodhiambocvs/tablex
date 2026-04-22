package core

// Table represents a booking table inside a club
type Table struct {
	ID        int
	ClubID    int
	Type      string // VVIP, VIP, REGULAR
	Price     int
	Available bool
}