package main

import "fmt"

func main() {
	lnum := 70
	ltrue := (lnum == 73)
	lge := (lnum <= 73)
	lse := (lnum >= 73)
	lne := (lnum != 70)
	lst := (lnum < 70)
	lgt := (lnum > 70)

	fmt.Println( ltrue, lge, lse, lne, lst, lgt)

}
