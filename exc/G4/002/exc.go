package main

import "fmt"

func main() {
s_int := []int{ 1, 2 ,3, 4, 5, 6, 7, 8, 9, 0}
 for ind, v_int := range s_int{
	 fmt.Println(ind, v_int)
 }
   fmt.Printf("%T\n%v", s_int, cap(s_int))
}