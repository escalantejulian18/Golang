package dummy

/*
	Paquete locales.
	como  este archivo no se va a ejecutar no lleva una funcion main.
	todas las funciones que empiecen con minuscula son privadas.
	init es una funcion que contiene todas las configuraciones 
	para que el paquete pueda funcionar
*/

var hola_mundo string

func init(){
	hola_mundo = "¡Hola Mundo Go!"
}

func HolaMundo() string{
	return hola_mundo
}