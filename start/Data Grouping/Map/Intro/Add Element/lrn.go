package main

import "fmt"

func main() {
	//Set map fields
	mpnum := map[int]string{
		001: "one",
		002: "two",
		003: "three",
		004: "four",
		005: "five",
	}

	fmt.Println(mpnum)

	//Add value to mapped table
	mpnum[006] = "six"

	fmt.Println(mpnum[001])
	// check if value in map table
	validate, ok := mpnum[006]
	fmt.Println("not ok \t", validate)
	fmt.Println(ok)

	// does check and print if exists
	if validate, ok := mpnum[002]; ok {
		fmt.Println("if in map table, value =", validate)
	}

	for key, val := range mpnum {
		fmt.Println(key, val)
	}

}
