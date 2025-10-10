package main

import "fmt"

func main() {
	// Your code goes here
	fmt.Println("Hello, World!")
	var arr [5]int
	fmt.Println("Array:", arr)
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("Updated Array:", arr)
	fmt.Println("Length of Array:", len(arr))

	arr2 := [3]string{"Go", "Python", "Java"}
	fmt.Println("String Array:", arr2)
	fmt.Println("Length of String Array:", len(arr2))

	var twodimarr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Array:", twodimarr)

	var slice2 []int = arr[1:4]
	fmt.Println("Slice from Array:", slice2)
	fmt.Println("Length of Slice:", len(slice2))
	fmt.Println("Capacity of Slice:", cap(slice2))
	slice2 = append(slice2, 60, 70)
	fmt.Println("Slice after append:", slice2)
	fmt.Println("Length of Slice after append:", len(slice2))
	fmt.Println("Capacity of Slice after append:", cap(slice2))

	slice3 := make([]string, 2, 5)
	slice3[0] = "Hello"
	slice3[1] = "World"
	fmt.Println("Slice created with make:", slice3)
	fmt.Println("Length of Slice created with make:", len(slice3))
	fmt.Println("Capacity of Slice created with make:", cap(slice3))
	slice3 = append(slice3, "from", "Go")
	fmt.Println("Slice after append:", slice3)
	fmt.Println("Length of Slice after append:", len(slice3))
	fmt.Println("Capacity of Slice after append:", cap(slice3))

	var copyslice []int = make([]int, len(arr))
	copy(copyslice, arr[:])
	fmt.Println("Copied Slice:", copyslice)
}
