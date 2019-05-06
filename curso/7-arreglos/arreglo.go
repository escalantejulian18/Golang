/*
	arreglos: cantidad de elementos y tipo de esos datos
	Son estaticos

*/

package main

import "fmt"

func main(){

	var arreglo [10] int // Declaración 1

	arreglo2 := [10] int{1,2,3} //Declaración 2
	
	longitudArreglo := len(arreglo2)	

	fmt.Println(arreglo, arreglo2, longitudArreglo)
}

