package main

import "fmt"

type usuario struct {
	nome     string
	idate    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var usuario1 usuario
	usuario1.nome = "Teste"
	usuario1.idate = 12
	fmt.Println(usuario1)

	endereco1 := endereco{"Endereco qualquer", 0}

	usuario2 := usuario{"Teste2", 21, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Teste3"}
	fmt.Println(usuario3)

}
