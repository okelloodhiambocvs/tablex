package core

// Booking represents a table reservation
type Booking struct {
	ID         int
	UserName   string
	TableID    int
	People     int
	Budget     int
	Status     string // CONFIRMED or PENDING
}