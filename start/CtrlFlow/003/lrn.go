package main

import "fmt"

func main() {
	x := 1
	// single cond, without itteration, Change managed in run block
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
