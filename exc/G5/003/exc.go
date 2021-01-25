package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type sedan struct {
	vehicle
	luxury bool
}
type truck struct {
	vehicle
	fourx bool
}

func main() {
	st0sedan := sedan{
		luxury: true,
		vehicle: vehicle{
			doors:  4,
			colour: "blue",
		},
	}
	st0fwd := truck{
		fourx: false,
		vehicle: vehicle{
			doors: 2,
			colour: "red",
		},
	}
	fmt.Println(st0fwd, st0sedan)
}
