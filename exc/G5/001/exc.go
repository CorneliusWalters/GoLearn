package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	st0pers := person{
		first:    "Corne",
		last:     "Walters",
		icecream: []string{"toffee", "coffee"},
	}
	st1pers := person{
		first:    "Doriana",
		last:     "Meuthen",
		icecream: []string{"chocolate", "strawberry"},
	}

	x1pers := []person{st0pers, st1pers}

	for _, vst0pers := range x1pers {
		fmt.Printf("first name: %v \nlast name: %v \nLikes these flavour icecream  \n", vst0pers.first, vst0pers.last)
		for _, vst0flav := range vst0pers.icecream {
			fmt.Println("\t\t", vst0flav)

		}

	}
}
