package main

import "fmt"

func main() {
	s0int := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for ind, v0int := range s0int {
		fmt.Println(ind, v0int)
	}
	fmt.Printf("%T\n%v", s0int, cap(s0int))
}
