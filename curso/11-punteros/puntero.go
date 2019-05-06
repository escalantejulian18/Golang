package main

import "fmt"

func main(){
	/*
		- puntero: es una dirección de memoria.
		- en vez del valor, tenemos el lugar donde esta el valor.
		- X, Y -> as123d --> 5
		- X -> as123d -> 6
		- Y ¿? -> 6
		- *T -> *int, *string
	}	- el valor cero de un puntero es nil
	*/

	var x,y *int
	entero := 5

	x = &entero
	y = &entero 

	fmt.Println(x,y)

	*x = 6
	
	fmt.Println(*x,*y)

}