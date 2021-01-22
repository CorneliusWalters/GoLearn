package main

import "fmt"

func main() {
	x0str := []string{"CH", "Walters", "Student"}
	x1str := []string{"T", "McCloud", "Lecturer"}

	xx0str := [][]string{x0str, x1str}
	fmt.Println(xx0str)

	for ind0, vxx0str := range xx0str {
		fmt.Println("Record Num", ind0)
		for ind1, vx0str := range vxx0str {
			fmt.Printf("\t Value: \t %v \t Index: %v \n", vx0str, ind1)

		}
	}
}
