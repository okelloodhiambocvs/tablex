package store

import "tablex/backend/core"

var splits []core.SplitRequest
var splitID = 1

// SaveSplit stores a split request
func SaveSplit(s core.SplitRequest) core.SplitRequest {
	s.ID = splitID
	splitID++

	splits = append(splits, s)
	return s
}

// GetSplits returns all split requests
func GetSplits() []core.SplitRequest {
	return splits
}