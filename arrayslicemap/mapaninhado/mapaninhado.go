package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15000.00,
			"Guga Pereira":   12000.00,
		},
		"J": {
			"José João": 6000.00,
		},
		"P": {
			"Pedro Junior": 15000.00,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
