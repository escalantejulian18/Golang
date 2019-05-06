package main

import(
	"fmt"
	"strings"
	"time"
)

func main(){
	// Código síncrono, respeta el orden 	
	mi_nombre_lento("Julián")
	fmt.Println("que aburrido es esperar")
}

func mi_nombre_lento(nombre string){

	// la función Split separa la cadena en mas cadenas
	letras := strings.Split(nombre,"")

	for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}