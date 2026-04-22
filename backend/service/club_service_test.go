package service

import "testing"

func TestGetAllClubs(t *testing.T) {
	clubs := GetAllClubs()

	if clubs == nil {
		t.Error("expected clubs but got nil")
	}
}