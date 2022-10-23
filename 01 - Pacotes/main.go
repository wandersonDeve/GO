package main

import (
	"fmt"
	"modulos/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("com")
	fmt.Println(erro)
}