package main

import "fmt"

func main() {
	f1()
	f2("f2")
	ps0 := f3("f3")
	f2(ps0)
	ps1, pb0 := f4("f3", "return" ) //seperate with comma on more than one returning paramaters
	if pb0 == true {
		f2(ps1)		
	}
}

// func ( r reciever ) identifiers(Par) (Return(s)){ ...}

func f1() {
	fmt.Println("f1")
}

func f2(ar0str string) {
	fmt.Println("call Function ", ar0str)
}
func f3(ar1str string) string {
	return fmt.Sprint("return ", ar1str)
}
func f4(ar2str string, ar3str string) ( string, bool){
	rs1 := fmt.Sprint("F4 check value ", ar2str)
	rb2 := true
	return rs1, rb2
}
