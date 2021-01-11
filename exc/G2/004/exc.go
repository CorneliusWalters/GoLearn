package main

import "fmt"

func main() {

	vint := 123

	fmt.Printf( "%d\t%b\t%x\n", vint, vint, vint )
	
	vbin := vint << 1
	fmt.Printf( "%d\t%b\t%x\n", vbin, vbin, vbin )
	
}
