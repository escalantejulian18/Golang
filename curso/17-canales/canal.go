package main

import "fmt"

func main(){

	// Creamos un canal donde se pueden transportar cadenas
	canal := make(chan string)

	go func(canal chan string){
		for{
			var nombre string
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			canal <- nombre //enviamos el nombre por el canal
		}
	}(canal)

	msg := <- canal //recibimos la informacion del canal

	fmt.Println("información recibida: " + msg)
}