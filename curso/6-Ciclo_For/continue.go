package main

import "fmt"

func main(){
	
	i := 0
	for {
		if i == 2{
			i++
			continue
		}

		fmt.Println("Hello ciclo While en Go", i)
		i++
		
		if i > 10{
			break // finaliza el ciclo
		}
	}
}