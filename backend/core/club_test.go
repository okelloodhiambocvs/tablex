package core

import "testing"

func TestClubStructure(t *testing.T) {
	club := Club{
		ID:   1,
		Name: "Test Club",
	}

	if club.Name != "Test Club" {
		t.Error("Club name mismatch")
	}
}