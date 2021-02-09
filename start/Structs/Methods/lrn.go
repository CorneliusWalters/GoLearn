package main

import "fmt"
type ty0str struct{
	first string
	last string
}

type ty1str struct{
	ty0str
	company string
}

func (m0 ty1str) f0str(){
	fmt.Println("Employee Details")

}

func main() {

   st0emp := ty1str{
	ty0str:  ty0str{
		   first: "Cornelius",
		   last: "Walters",
	   },
	   company: "AVI",
   }
   fmt.Println(st0emp)
   st0emp.f0str()

}