package main

import "fmt"

func main(){

	slice := []int{1,2,3,4}
	copia := make([]int,4)
	/*
		copy(destino, fuente)

		La funcion copia el munimo de elementos entre ambas estructuras
		es decir que el arr destino debe tener la misma longitud que el 
		arr fuente para que se cipie completo.

		copia := make([]int,len(arr_fuente),cap(arr_fuente)*2)
	*/
	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)

}