package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por *
func inc2(n *int) {
	// revisão:* é para desreferenciar, ou seja, ter acesso no valor ao qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisão: & é usado para obter o endereço da variável
	inc2(&n) // por referência
	fmt.Println(n)
}
