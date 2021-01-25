package main

import "fmt"

// declare type of structure
//
type person struct {
	first string
	last  string
	age   int
}

type address struct {
	street string
	numb   string
	suburb string
}

type details struct {
	person
	address
}

func main() {

	st0det := details{
		person:  person{
			first: "Cornelius",
			last:  "Walters",
			age:   33,
		},
		address: address{
			street: "Regent Road",
			numb:   "67",
			suburb: "Sea Point",
		},
	}
	st1det := details{
		person:  person{
			first: "Jerusha",
			last:  "Ramsamy",
			age:   43,
		},
		address: address{
			street: "Westwick",
			numb:   "11",
			suburb: "Gardens",
		},
	}

	// can access with dot notation
	fmt.Println(st0det.numb, st1det)

}
