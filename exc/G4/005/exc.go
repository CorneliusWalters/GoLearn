package main

import "fmt"

func main() {
	s0int := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}
	s0int = append(s0int[:3], s0int[6:]...)
	fmt.Println(s0int)
}
