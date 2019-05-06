/*
	Go es un lenguaje fuertemente tipado: Es importente declarar el tipo de la variable.
	Todas lad variables iniciadas tiene un valor "cero": 0 para entero, "" para cadena, true para bool.
	Todas las variables declaras deben ser usadas. (Da error de compilación)
*/

package main

import "fmt"

func main(){
	var x,y,z int
	var cadena string
	var bool bool

	const pi = 3.14

	x = 25 // Asignación.
	w := 35 // Tipado dinamico, interpreta el tipo del dato en tiempo de ejecución.

	nombre := "coco"

	fmt.Println(x, y, z, cadena, bool, pi, w, nombre)

}