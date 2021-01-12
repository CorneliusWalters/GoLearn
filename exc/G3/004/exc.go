package main

import "fmt"

func main() {
	for vrem := 10; vrem <= 100; vrem++ {
		fmt.Printf("remainder of %v div 4 eq: \t %v\n", vrem, vrem%4)
	}

}
