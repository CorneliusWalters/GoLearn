package main

import "fmt"

func main() {

	for vcode := 33; vcode <= 122; vcode++ {
		fmt.Printf("'%v': is the letter:\t %#U\t with hex code: '%#x'\n", vcode, vcode, vcode)
	}
}
