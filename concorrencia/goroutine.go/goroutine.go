package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("Pedro", "Só posso falar dps de vc", 1)

	// go fale("Maria", "Pq vc não fala comigo?", 3)
	// go fale("Pedro", "Só posso falar dps de vc", 1)

	go fale("Maria", "Entendi!", 10)
	fale("Pedro", "Parabéns", 5)
}
