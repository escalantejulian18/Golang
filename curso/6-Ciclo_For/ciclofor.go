/*
	En Go solo existe un solo ciclo: el ciclo For.
	en muchos lenguajes los tres terminos(inicio, condición e incremento) del for son fundamentales 
	pero para go esto es mas flexible.

*/

package main

import "fmt"

func main(){

	for i := 0 ; i < 10 ; i++{
		fmt.Println("Hello ciclo For en Go")
	}
	
	i := 0
	for  i < 10  {
		fmt.Println("Hello ciclo While en Go")
		i++
	}
}