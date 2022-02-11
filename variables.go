package main

import (
	"fmt"
)

var z int = 38 //como buena practica scoup  dentro del bloque de codigo y no general

func main() {
	edad := 37 // declaracion corta de variables con el operador :=   va dentro del cuerpo de la funcion {}
	nombre := "William Guillermo"
	fmt.Println(nombre, edad, "Vallarta")

	y := 2 // asignacion corta  pero una vez asignada en el programa se debe utilizar
	x := 3
	fmt.Println(x + y)
	fmt.Println(z)
}
