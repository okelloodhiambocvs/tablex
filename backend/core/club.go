package core

// Club represents a nightlife venue in Kisumu
type Club struct {
	ID       int
	Name     string
	Location string
	Tables   []Table
}