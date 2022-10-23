package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro, *ponteiro)

	variavel3 = 300
	fmt.Println(variavel3, ponteiro, *ponteiro)


}
