package postgresloader

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(5, 2)
	if result != 7 {
		t.Errorf("Expected 3, got %d", result)
	}
}
