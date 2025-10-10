package main

import (
	"fmt"
)

// Strut represents a basic structure for demonstration.
type Strut struct {
	ID    int
	Name  string
	Value float64
}

// NewStrut creates a new Strut instance.
func NewStrut(id int, name string, value float64) *Strut {
	return &Strut{
		ID:    id,
		Name:  name,
		Value: value,
	}
}

// SetValue sets the Value field of Strut.
func (s *Strut) SetValue(val float64) {
	s.Value = val
}

// Display prints the details of the Strut.
func (s *Strut) Display() {
	fmt.Printf("ID: %d, Name: %s, Value: %.2f\n", s.ID, s.Name, s.Value)
}

// UpdateName changes the Name field of Strut.
func (s *Strut) UpdateName(newName string) {
	s.Name = newName
}

func main() {
	s := NewStrut(1, "Example", 42.0)
	s.Display()
	s.SetValue(100.5)
	s.UpdateName("Updated Example")
	s.Display()
}
