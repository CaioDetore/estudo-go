package main

import "fmt"

func main() {
	n := 1

	switch {
	case n >= 9 && n <= 10:
		fmt.Println("A")
	case n >= 8 && n < 9:
		fmt.Println("B")
	case n >= 5 && n < 8:
		fmt.Println("C")
	case n >= 3 && n < 5:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
