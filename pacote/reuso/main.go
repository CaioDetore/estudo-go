package main

import (
	"estudo-go/mypkg"
	"fmt"
)

func main() {
	area := &mypkg.Area{}
	fmt.Println(area.Circ(6.0))
}
