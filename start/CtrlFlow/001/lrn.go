package main

import "fmt"

func main() {
	// var cond initialised  				[for *var*;]
	// condition evaluation  				[ l1 <= 10;]
	// itteration clause post statement		[ l1 ++]

	for l1 := 0; l1 <= 10; l1++ {

		fmt.Printf("loop l1 is: %d \n", l1)
		for l2 := 7; l2 >= 7 && l2 <= 9; l2++ {
			fmt.Printf(" loop l2 is: %d \t\n", l2)
		}
	}

	fmt.Println("hi")
}
