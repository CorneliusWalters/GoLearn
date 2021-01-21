package main

import "fmt"

func main() {
	a0int := [5]int{999, 100, 200, 300, 400}

	v0cap := cap(a0int)

	for ind := 0; ind < v0cap; ind++ {
		fmt.Println(a0int[ind])

	}
	fmt.Printf("%T\t capacity of array: %v", a0int, v0cap)

}
