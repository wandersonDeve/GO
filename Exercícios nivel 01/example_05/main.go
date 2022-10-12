package main

import "fmt"

type number int
var x number

var y int

func main() {
	fmt.Printf("valor de x = %v\t Tipo de x = %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Printf("valor de y = %v\t Tipo de y = %T\n", y, y)
}
