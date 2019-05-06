package main

import(
	"fmt"
	"strings"
	"time"
)

/*
	La palabra "go" indica que ese función se ejecutará en forma concurrente.
	El límite de rutinas es miles.

*/
func main(){
	// Código síncrono, respeta el orden 	
	go mi_nombre_lento("Julián")
	fmt.Println("que aburrido es esperar")

	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lento(nombre string){

	// la función Split separa la cadena en mas cadenas
	letras := strings.Split(nombre,"")

	for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}