/* 
	Al igual que las variables, los paquetes declarados deben ser usados.
	"Atoi" retorna dos valores: nuestro valor transformado, valor por si surge un error
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){

	edad := "22"
	edad_int,_ := strconv.Atoi(edad) // Operador "_" desecha todo otro valor

	fmt.Println(10 + edad_int)

}