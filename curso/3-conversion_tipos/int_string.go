/* 
	Al igual que las variables, los paquetes declarados deben ser usados.
	
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){

	edad := 22
	edad_str := strconv.Itoa(edad)

	fmt.Println("Mi edad es " + edad_str)

}