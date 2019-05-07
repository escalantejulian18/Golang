package main

import(
	"net/http"
	"encoding/json"
)

type Curso struct{
	Title string `json:"title"`
	NumeroVideos int `json:"numero_videos"`
}



func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		curso := Curso{"Curso de Go",26}
		json.NewEncoder(w).Encode(curso)
	})

	http.ListenAndServe(":8000",nil)

}