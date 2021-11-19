package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(1))
	fmt.Println(tipo(1.0))
	fmt.Println(tipo("1"))
	fmt.Println(tipo(func() {}))
}
