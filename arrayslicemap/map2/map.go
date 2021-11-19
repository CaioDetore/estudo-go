package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      13500.00,
		"Gabriela Silva": 14500.00,
		"Paulo Junior":   1200.00,
	}

	funcESalarios["Rafael Filho"] = 1400.00
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
