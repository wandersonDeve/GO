package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Arrays e Slices")

	var arr1 [5]string
	arr1[0] = "Posição 1"
	fmt.Println(arr1)

	arr2 := [5]string{"posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(arr2)

	arr3 := [...]int{1,2,3,4,5,6}
	fmt.Println(arr3)

	// SLICES

	slice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	slice = append(slice, 10)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	arr2[1] = "Posição alterada"
	fmt.Println(slice2)

}