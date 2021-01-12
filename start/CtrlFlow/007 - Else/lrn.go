package main

import "fmt"

func main() {
	var vcond int = 101
	if vcond == 100 {
		fmt.Println("vcond is not 100")
	} else if vcond == 102 {
		fmt.Println("vcond is not 102")
	} else {
		fmt.Println("vcond did not match any condition")
	}
}
