package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// número inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	il := math.MaxInt64
	fmt.Println("O valor máximo de int é", il)
	fmt.Println("O tipo de il é", reflect.TypeOf(il))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode(int32)
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Texto string"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String com multiplas linhas
	s2 := `Texto
	String`
	fmt.Println("O tamanho da string é", len(s2))

	// char?
	// var a char = 'a'
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
