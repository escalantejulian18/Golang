package main

import "fmt"

func main(){

	slice := make([]int,3,4)
	slice = append(slice,2) // agrega al final del slice
	//slice = append(slice,2)
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}