package postgresloader

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := addition(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
