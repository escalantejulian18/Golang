package main

import (
	"io/ioutil"
	"fmt"
)
/*
	Para ejecutar debemos estar 
	posicionados en la carpeta 
	del archivo go.

	Err que puede pasar es que no
	encuentre el archivo a leer

*/

func main(){

	archivo, err := ioutil.ReadFile("./hola.txt")
	 
	if err != nil{
		fmt.Println("Hubo un error")
	}

	fmt.Println(string(archivo))
}