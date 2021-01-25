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
	m0Pers := make(map[string]person, 0)

	for _, vst0pers := range x1pers {
		m0Pers[vst0pers.last] = vst0pers
	}
	// fmt.Println(m0Pers)
	for km0, vmp0pers := range m0Pers {
		fmt.Println(km0)
			for _, vst0flav := range vmp0pers.icecream {
				fmt.Println("\t\t", vst0flav)
			}
	}
}
