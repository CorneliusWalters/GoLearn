package main

import "fmt"

func main() {

   vslice1 := make([]int, 15, 16)
   fmt.Println(vslice1)
   fmt.Println(len(vslice1))
   fmt.Println(cap(vslice1))

}