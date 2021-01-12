package main

import "fmt"

func main() {

	for vcnt := 0; vcnt < 100; vcnt++ {
		if vcnt%3 == 0 {
			fmt.Printf("%v is divisible by 3\n", vcnt)
		}
	}

}
