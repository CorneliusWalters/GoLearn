package main

import "fmt"

func main() {

	x0int := []int{ 1,2,3,4,5,6,7,8,9,0};
	fs0int :=sum(x0int...)

   fmt.Println(fs0int)

}

func sum( ar0int ...int ) int{
	v1sum :=0
	for _, v01int := range ar0int {
        v1sum += v01int	
	}
	return v1sum
}