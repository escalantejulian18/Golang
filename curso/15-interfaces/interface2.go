package main

import "fmt"

type User interface{
	Permisos() int // 1 - 5
	NombreUsuario() string
}

// Administrador
type Admin struct{
	nombre string
}

func (this Admin) Permisos() int{
	return 5
}
func (this Admin) NombreUsuario() string{
	return this.nombre
}


// Editor
type Editor struct{
	nombre string
}

func (this Editor) Permisos() int{
	return 3
}
func (this Editor) NombreUsuario() string{
	return this.nombre
}
 
func auth(user User) string{
	permisos := user.Permisos()
	if permisos > 4{
		return user.NombreUsuario() + ": Tiene permisos de Administrador"
	}else if permisos == 3{
		return user.NombreUsuario() + ": Tiene permisos de Editor"
	}
	return user.NombreUsuario() + ": No tiene permisos asignados"
}

func main(){

	admin := Admin{"Julián"}
	editor := Editor{"Fulano"}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))

}