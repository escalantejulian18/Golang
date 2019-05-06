package main

import "fmt"

/*
	- nos permite replicar la Herencia y la sobre-escritura.
	- nos deja usar campos de otras estructuras tambien declaras.
*/


type Human struct{
	nombre string
}

func (this Human) hablar() string{
	return "bla bla bla..."
}

type Dummy struct{
	nombre string
}

type Tutor struct{
	Human
	Dummy
}

func (this Tutor) hablar() string{
	this.Human.hablar()
	return "Hola" + " " + this.Human.hablar()
}

func main(){

	tutor := Tutor{Human{"Julián"}, Dummy{"Julián"}}
	fmt.Println(tutor.Human.nombre)
	fmt.Println(tutor.Dummy.nombre)
	fmt.Println(tutor.hablar())


}