package main

import "fmt"

func main() {

   vslice1 := []int{1, 2, 3, 4, 5}
   fmt.Println(vslice1)
   vslice1 = append(vslice1, 6, 7, 8, 9)
   fmt.Println(vslice1)
   vslice2 := []int{10, 11, 12, 13}
   vslice1 = append(vslice1, vslice2...)
}