/*
	Se basan en los arreglos, pero diamicos.
	Se puede redimencionar.
	Un Slice sin datos iniciales es igul a nulo.
	Elementos de un slice:
		1. Longutud del arreglo al que referencia.
		2. Puntero al arreglo.
		3. Capacidad.
*/

package main

import "fmt"

func main(){

	var slice [] int // Declaración 1
	slice2 := [] int{1,2,3,4}

	fmt.Println(slice,slice2)
}
