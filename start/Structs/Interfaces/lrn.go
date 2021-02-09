package main

import "fmt"

type identity struct{
	first string
	last  string
	pref  string
	idno    int
}
type address struct{
	street string
	stnumber int
	building string
	blnumber int
	suburb string
}
type person struct{
	identity
	address
	real bool
}

func (pers person)
func main() {

   fmt.Println("hi")

}