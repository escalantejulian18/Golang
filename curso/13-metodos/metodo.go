package main

import "fmt"

/*
	a pesar que Go no es un lenguaje que permita POO, 
	permite crear metodos para las estructuras.
*/

type User struct{
	edad int
	nombre, apellido string
}

func (usuario *User) nombre_completo() string{
	return usuario.nombre + " " + usuario.apellido 
}

func (this *User) set_name(nom string){
	this.nombre = nom
}


/*
	otra forma de declarar, familiarizandolo con POO
	func (this User) nombre_completo() string{
		return this.nombre + " " + this.apellido 
	}
*/

func main(){

	usuario := new(User)
	usuario.nombre = "Julián"
	usuario.apellido = "Esc"

	usuario.set_name("Jul")

	fmt.Println(usuario.nombre_completo())

}