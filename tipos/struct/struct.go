package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método com receiver (receptor)
func (p produto) valorComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.valorComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.valorComDesconto())
}
