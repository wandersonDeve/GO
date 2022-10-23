package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main(){
	soma := soma(10,20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado, "<- Resultado")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10,15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	reultSoma, _ := calculosMatematicos(10,15)
	fmt.Println(reultSoma)

	_, resultSubtracao := calculosMatematicos(10,15)
	fmt.Println(resultSubtracao)
}