package main

import "fmt" 

/*
	== igual a
	!= diferente de
	>  mayor que
	<  menor que
	>= mayor o igual que 
	<= menor o igual que
	&& and
 	|| or

*/

func main(){

	x := 10
	y := 12

	if x >= y {
		fmt.Println("Hola Mundo")
		fmt.Printf("%d es mayor que %d\n",x,y)

	}else{
		fmt.Printf("%d es menor que %d\n",x,y)
	}
}