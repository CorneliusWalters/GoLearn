package main

import "fmt"

func main() {
	for l1 := 0; l1 <= 10; l1++ {
		for l2 := 0; l2 > 7; l2++ {
			fmt.Printf("loop l1 is: %d \t loop l2 is: %d \n", l1, l2)
		}
	}
}
