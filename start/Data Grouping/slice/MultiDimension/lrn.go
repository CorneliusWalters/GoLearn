package main

import "fmt"

func main() {

	sdim1I := []string{"entry1_1", "entry1_2", "entry1_3"}
	sdim1II := []string{"entry2_1", "entry2_2", "entry2_3"}

	sdim2I := [][]string{sdim1I, sdim1II}

	fmt.Println(sdim2I)

}
