package main

import (
	"fmt"
	"unicode"
)

func main() {
	var irune rune
	for irune = 33; irune < 122; irune++ {
		srune := fmt.Sprintf("\t%q", irune)
		switch unicode.IsUpper(irune) {
		case true:

			fmt.Printf("%v\n", irune)
			for vcnt := 0; vcnt <= 2; vcnt++ {
				println(srune)
			}
		}
	}

}
