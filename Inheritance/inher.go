package main

import "fmt"

// Base struct
type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println("Animal speaks")
}

// Derived struct embedding Animal (inheritance-like)
type Dog struct {
	Animal // Embedding
	Breed  string
}

func (d *Dog) Speak() {
	fmt.Printf("%s the %s barks\n", d.Name, d.Breed)
}

func main() {
	a := Animal{Name: "Generic"}
	a.Speak()

	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}
	d.Speak()
}
