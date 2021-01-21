package main

import "fmt"

func main() {
	s0int := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(s0int[:4])
	fmt.Println(s0int[5:])
	fmt.Println(s0int[3:7])
	fmt.Println(s0int[2:6])
}
