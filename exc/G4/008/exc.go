package main

import "fmt"

func main() {
	m0str := map[string][]string{
		"C_Walters": []string{"Learning GoLang", "udemy", "Lazy AF", "Age"},
		"Z_Kader":   []string{"Integration Manager", "AVI", "SAP PI", "Vape"},
		"S_Meyer":   []string{"Mobile Manager", "AVI", "Lesbian", "Trupple"},
		"A_Jordi":   []string{"Exec", "AVI Systems", "Moving", "I will miss her"},
	}
// on map index replaced with key, 
	for key0, vm0 := range m0str {
		fmt.Println(key0, "Details:")
		for ind0, vx0str := range vm0 {
			fmt.Println("\t", ind0, vx0str)
			
		}
	}
}
