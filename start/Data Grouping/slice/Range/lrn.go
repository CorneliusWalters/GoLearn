package main

import "fmt"

func main() {
	// vslice := []int{}  - Composite Lit
	vslice := []int{4, 5, 3, 7, 8}

	for ind, islice := range vslice {
		fmt.Println(vslice, ind, islice)

	}

}
