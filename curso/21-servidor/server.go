package main
 
import (
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		//http.ServeFile(w,r,"index.html")
		http.ServeFile(w,r,r.URL.Path[1:])
		/*
			extrae todos los objetos a través Request
		*/
	})
	http.ListenAndServe(":8000",nil)
}