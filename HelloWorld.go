package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func add(a, b int) (x, y int) {
	x = a
	y = b
	return
}

func sub(a, b int) int {
	return a - b
}

func fib(n int, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fiba := make([]int, n)
	fiba[0] = 0
	fiba[1] = 1
	for i := 2; i < n; i++ {
		fiba[i] = fiba[i-1] + fiba[i-2]
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Before data written in channel")
	ch <- fiba
	fmt.Println("data written in channel")

	//fmt.Println(fiba)
}

func main() {
	num1 := 5
	num2 := 7
	//m := add(num1, num2)
	//t.Printf("Sum of %d and %d is %d\n", num1, num2, sum)
	fmt.Println(add(num1, num2))
	fmt.Println(math.Pi)

	fmt.Println(sub(num1, num2))
	const myconst = 1
	fmt.Printf("type : %T, value :%v", myconst+1.1, myconst+1.1)

	const (
		a = iota + 2 // 0
		b
		c
	)

	fmt.Println(a, b, c) // 0 1 2

	var wg sync.WaitGroup
	ch := make(chan []int, 5)
	wg.Add(1)
	go fib(25, ch, &wg)
	//fmt.Println(<-ch)

	fmt.Println("waiting ")

	wg.Add(1)
	go fib(5, ch, &wg)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	wg.Wait()
	//time.Sleep(time.Millisecond * 1500)

	mm := make(map[string]int)
	mm["one"] = 1
	mm["two"] = 2
	fmt.Println(mm)
}

func disp() {
	fmt.Println("Hello")
}
