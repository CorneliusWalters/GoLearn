package main

import "fmt"

func main() {
	a_int := [5]int{999, 100, 200, 300, 400}

	v_cap := cap(a_int)

	for ind := 0; ind < v_cap; ind++ {
		fmt.Println(a_int[ind])

	}
	fmt.Printf("%T\t capacity of array %v", a_int, v_cap)

}
