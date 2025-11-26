package main

import (
	"fmt"
	"reflect"
)

// Strut represents a basic structure for demonstration.
type Strut struct {
	ID    int
	Name  string
	Value float64
}

type Account struct {
	AccountNum int
	Balance    int
	p          Personal
}

type Personal struct {
	Name  string
	Email string
	Phone int
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

func (a *Account) DisplayAccInfo() {
	fmt.Println("From method", a.AccountNum, a.Balance, a.p.Name)

}

func main() {
	s := NewStrut(1, "Example", 42.0)
	s.Display()
	s.SetValue(100.5)
	s.UpdateName("Updated Example")
	s.Display()

	var a Account
	a.AccountNum = 20
	a.Balance = 100
	fmt.Println(a)

	var aa = Account{AccountNum: 10, p: Personal{Name: "Sen"}}
	fmt.Println(aa)
	aa.DisplayAccInfo()

	fmt.Println(" TYpe of acc", reflect.TypeOf(aa))
	fmt.Println(" Value of acc", reflect.ValueOf(aa))

	var aaa = Account{}
	fmt.Println(aaa)

	if aa == aaa {
		fmt.Println("aa == aaa")
	} else {
		fmt.Println("aa != aaa")
	}

	var aaaa = new(Account)
	aaaa.AccountNum = 324
	aaaa.Balance = 1000
	fmt.Println(*aaaa)
}
