package main

import "fmt"

func main() {

	if true {
		println("001")
	}

	if false {
		println("002")
	}
	if !true {
		println("003")
	}
	if !false {
		println("004")
	}
// 
	if x := 100; x == 99 {
		println("005")
	}
}