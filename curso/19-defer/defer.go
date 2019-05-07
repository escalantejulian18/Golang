package  main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){

	ejecicion := readFile()
	fmt.Println(ejecicion)
}

func readFile() bool {
	archivo, err := os.Open("./hola.txt")
	defer func(){
		archivo.Close()
		fmt.Println("Defer")
	}()

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
	return true

}