package main

import "fmt"

func main() {
	s0int := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}
	s0int = append(s0int, 10)
	fmt.Println(s0int)
	s0int = append(s0int, 11, 12, 13)
	fmt.Println(s0int)
	s1int :=[]int{14, 15, 16, 17}
	s0int = append(s0int, s1int[:]...)
	fmt.Println(s0int)

}
