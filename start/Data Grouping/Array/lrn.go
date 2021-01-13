package main
import "fmt"
func main() {
	var varray [10]int
	fmt.Println(varray)
	varray[5] = 6
	fmt.Println(varray)
	fmt.Println(len(varray))
	
}