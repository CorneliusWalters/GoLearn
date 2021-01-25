package main

import "fmt"

func main() {
	st0anon := struct {
		company string
		employee map[int][]string
		application []string
	}{
		company: "AVI",
		employee: map[int][]string{
			1234: {"Cornelius", "Walters", "Sea Point"},
			2135: {"Adam", "Van Tonder","Hope Field"},
			5421: {"Shireen", "Meyer", "Thrupple"},
		},
		application: []string{"SAP", "MSBI", "Office"},
	}
	
	fmt.Println(st0anon.company, st0anon.application)
}
