package bites1

import "testing"

func TestBites1(t *testing.T) {
	result := bobsAge()
	if result != 86 {
		t.Errorf("Result was incorrect, got %d, want: %d", result, 86)
	}
}
