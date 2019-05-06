package main

import "fmt"

/*
	- nos permite replicar la Herencia.
	- nos deja usar campos de otras estructuras tambien declaras.
*/

type Human struct{
	nombre string
}

type Tutor struct{
	Human
}

func main(){

	tutor := Tutor{Human{"Julián"}}

	fmt.Println(tutor.nombre)
}