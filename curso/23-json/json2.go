package main

import(
	"net/http"
	"encoding/json"
)

type Curso struct{
	Title string `json:"title"`
	NumeroVideos int `json:"numero_videos"`
}


type Cursos []Curso

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		cursos := Cursos{
			{"Curso de Go",26},
			{"Curso de Ruby",34},
			{"Curso de Python",30},
			{"Curso de NodeJs",25},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8000",nil)

}