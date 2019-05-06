package main

import "fmt"

type User interface{
	Permisos() int // 1 - 5
	NombreUsuario() string
}

type Admin struct{
	nombre string
}

func (this Admin) Permisos() int{
	return 5
}
func (this Admin) NombreUsuario() string{
	return this.nombre
}
 
func auth(user User) string{
	if user.Permisos() > 4{
		return user.NombreUsuario() + ": Tiene persomisos de Administrador"
	}
	return ""
}

func main(){

	admin := Admin{"Julián"}

	fmt.Println(auth(admin))
}