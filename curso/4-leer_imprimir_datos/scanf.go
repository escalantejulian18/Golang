package main

/*
	fmt es un paquete para lectura e impresión
*/

import ("fmt")

func main(){

	var edad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d",&edad)

	fmt.Printf("Mi edad es %d", edad)

}