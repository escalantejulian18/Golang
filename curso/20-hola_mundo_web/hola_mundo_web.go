package main

import(
	"net/http"
	"io"
	"fmt"
)

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/hola",func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "¡Hola!")	
	})	
	http.ListenAndServe(":8000",nil) // rooting defecto

}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "¡Hola Mundo Go!")	
}