package main

import "fmt"

type number int
var x number

func main() {
	fmt.Printf("%v\t%T\n", x, x)

	x = 42

	fmt.Println(x)
}
