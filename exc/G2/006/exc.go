package main

import "fmt"

const (
	cyr1 = 2020 + iota
	cyr2 = 2020 + iota
	cyr3 = 2020 + iota
	cyr4 = 2020 + iota
)

func main() {
	fmt.Println(cyr1, cyr2, cyr3, cyr4)

}
