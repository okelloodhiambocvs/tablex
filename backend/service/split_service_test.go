package service

import "testing"

func TestSplitSuggestion(t *testing.T) {

	split := CreateSplitSuggestion(1, "John", "Mary", 10000)

	if split.AmountA != 5000 {
		t.Error("Expected equal split amount")
	}

	if split.Status != "PENDING" {
		t.Error("Expected PENDING status")
	}
}