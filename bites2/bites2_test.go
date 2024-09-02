package main

import "testing"

func TestBites2(t *testing.T) {
	a, b, c := bites2()
	if a != 29 {
		t.Errorf("Result was incorrect, got %d, want: %d", a, 86)
	}
	if b != 400.68 {
		t.Errorf("Result was incorrect, got %f, want: %f", b, 400.68)
	}
	if c != "I love data types!" {
		t.Errorf("Result was incorrect, got %s, want: %s", c, "string")
	}
}
