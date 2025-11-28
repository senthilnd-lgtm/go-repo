package main

import (
	"fmt"
	"reflect"
)

func main() {
	sl := []string{"apple", "Orange"}
	sl = append(sl, "banana")
	fmt.Println(sl, len(sl), cap(sl))

	// define an array with 7 elements
	numbers := [7]int{0, 1, 2, 5, 798, 43, 78}
	fmt.Println("array value:", numbers)
	fmt.Println("array type:", reflect.TypeOf(numbers))

	// define a slice s based on the numbers array
	s := numbers[0:4]
	fmt.Println("slice value:", s)
	fmt.Println("slice type:", reflect.TypeOf(s))

	for _, val := range sl {
		fmt.Println(val)
	}

	ss := make([]int, 10)

	fmt.Println(ss)

	var sa = []int{10, 11, 14}
	fmt.Println("Slice s: ", sa)

	// create a destination slice
	c := make([]int, len(sa))

	// copy everything in s to c
	num := copy(c, sa) // returns the minimum number of elements in the slices
	fmt.Println("Number of elements copied:", num)
	fmt.Println("Slice c:", c)

	var mySlice = new([10]int)[0:5]
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3
	mySlice[3] = 4
	mySlice[4] = 5

	fmt.Println(mySlice)
	newSlice := RemoveIndex(mySlice, 2)
	fmt.Println(newSlice)

}

func RemoveIndex(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
