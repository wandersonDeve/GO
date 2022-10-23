package main

import "fmt"

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// ALIAS
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS DECIMAIS
	var numeroDecimal1 float32 = 123.45
	fmt.Println(numeroDecimal1)

	var numeroDecimal2 float64 = 12346952.25
	fmt.Println(numeroDecimal2)

	numeroDecimal3 := 123456.65
	fmt.Println(numeroDecimal3)

	// STRINGS
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// VALOR ZERO

	var texto string
	fmt.Println(texto, " <- Vazio")

	var number int64
	fmt.Println(number)

	
}
