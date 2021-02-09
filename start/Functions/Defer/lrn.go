package main

import "fmt"

func main() {
//  defer delays the executon of the function to the end of the current function,
// when the containing function exits
	defer pr0def( "Definition 1") 

	v1str :=  pr1def( "Definition 2")

	fmt.Println(v1str)
}

func pr0def( p0str string ) string{

	ret0str := fmt.Sprint(p0str)
	return ret0str
}

func pr1def( string ) string{

	ret0str := fmt.Sprint("defer2")
	return ret0str
}