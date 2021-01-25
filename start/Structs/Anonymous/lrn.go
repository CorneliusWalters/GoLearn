package main

import "fmt"

func main() {
// annonymous structures declare within the definition.
	st0det := struct{
			first string
			last  string
			age   int
		}{
			first: "Cornelius",
			last:  "Walters",
			age:   33,
		}
	fmt.Println(st0det.first)
	}

