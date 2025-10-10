package calc_test

import (
	calc "calctest2"
	"testing"
)

func TestAdd(t *testing.T) {
	result := calc.Add(3, 5)
	expected := 8
	if result != expected {
		t.Fatalf("Add(3,5) = %d; want %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := calc.Sub(10, 4)
	expected := 6
	if result != expected {
		t.Fatalf("Sub(10,4) = %d; want %d", result, expected)
	}
}
