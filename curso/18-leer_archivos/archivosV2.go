package  main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	archivo, err := os.Open("./hola.txt")

	if err != nil{
		fmt.Println("Hubo un error")
	}

	escaner := bufio.NewScanner(archivo)
	var i int
	for escaner.Scan(){
		i++
		linea := escaner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	archivo.Close()
}
