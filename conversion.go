package main

import (
	"fmt"
	"strconv" //paquete de conversion de tipos de variables
)

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	edad := 37
	edad_str := strconv.Itoa(edad) //funcion para convertir  INT  astring en Go
	fmt.Println("Mi eadad es:" + edad_str)

	// convertir un string a un entero
	peso := "58"
	peso_int, _ := strconv.Atoi(peso) //la conversion strcon retorna mas de una variable en
	//tal caso  se coloca un signo _  despues de la variuable para desechar los demas valores
	fmt.Println(peso_int + 10)

}
