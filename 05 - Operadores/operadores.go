package main

import "fmt"

func main() {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS
	fmt.Println(true && true)

	fmt.Println(true || false)
	fmt.Println(false || true)

	fmt.Println(!true)
	fmt.Println(!false)

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero -= 10
	fmt.Println(numero)


}