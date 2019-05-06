package main

import "fmt"

func main(){
	
	i := 0
	for  {
		fmt.Println("Hello ciclo While en Go")
		i++
		if i == 10{
			break // finaliza el ciclo
		}
	}
}