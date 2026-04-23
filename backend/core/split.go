package core

// SplitRequest represents a shared table request
type SplitRequest struct {
	ID        int
	TableID   int
	UserA     string
	UserB     string
	AmountA   int
	AmountB   int
	Status    string // PENDING, ACCEPTED, REJECTED
}