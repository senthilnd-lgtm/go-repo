package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// TODO: Add your test cases here
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Add(3, 5) = %d; want %d", result, expected)
	}
}
