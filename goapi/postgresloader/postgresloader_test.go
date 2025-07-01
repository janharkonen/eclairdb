package postgresloader

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := addition(5, 2)
	if result != 7 {
		t.Errorf("Expected 3, got %d", result)
	}
}
func TestLoadData(t *testing.T) {
	err := LoadData()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
