package main

import "fmt"

type foo int
var bar foo
const foobar = "200"

func main() {
	bar1 := int(foobar)
	bar = int(42)

   fmt.Println(bar)

}