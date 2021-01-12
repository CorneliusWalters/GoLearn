package main

import "fmt"

var chk bool

func main() {

	cnt := 0
	for chk == false {
		if cnt < 9 {
			cnt++
			fmt.Println(cnt)
		} else {

			chk = true
		}

	}

}
