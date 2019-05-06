package main

import "fmt"

func main(){

	arreglo := [3] int{1,2,3}
	slice := arreglo[0:3] // Slicing
	slice1 := arreglo[0:2] // de la posición 0 a la 2
	slice2 := arreglo[1:2]
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)

}