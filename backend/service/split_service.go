package service

import (
	"tablex/backend/core"
	"tablex/backend/store"
)

// CreateSplitSuggestion creates a match between two users
func CreateSplitSuggestion(tableID int, userA string, userB string, price int) core.SplitRequest {

	half := price / 2

	split := core.SplitRequest{
		TableID: tableID,
		UserA:   userA,
		UserB:   userB,
		AmountA: half,
		AmountB: half,
		Status:  "PENDING",
	}

	return store.SaveSplit(split)
}