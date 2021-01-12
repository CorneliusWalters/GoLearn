package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("case false")
	case 2 == 4:
		fmt.Println("case not equal")
	case 3 == 3:
		fmt.Println("EQUAL STATEMENT")
		fallthrough //evaluates next case
	case 5 == 5:
		fmt.Println("case not equal")
	default:
		fmt.Println("Default statement")
	}
}
