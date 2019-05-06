package main

import "fmt"

/*
	type define un nuevo tipo.
	struct define la estructura.
	las estructuras son mutables.
*/

type User struct{
	edad int
	nombre, apellido string
}

func main(){

	// var julian User // Defición 1
	// julian := User{25,"",""} // Definición 2
	
	julian := User{nombre: "Julian", apellido:"Esc", edad:25} // Definión 3

	fmt.Println(julian)
	

	julian2 := new(User) // Definición 4 -> es un puntero
	julian2.nombre = "Julián"
	fmt.Println((*julian2).nombre)
}