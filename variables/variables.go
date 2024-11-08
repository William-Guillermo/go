package main

import "fmt"

func main() {
	fmt.Println("variables en Go")
	var name string
	var lastname string
	var edad int

	name = "William"
	lastname = "Guillermo de La Paz"
	edad = 40
	juegoFavorito := "Futbol"

	fmt.Printf("Mi nombre es: %s\n ", name)
	fmt.Printf("Mi apellido es %s\n", lastname)
	fmt.Printf("Tengo %d years old :) \n", edad)
	fmt.Printf("Mi juego Favorito es %s \n", juegoFavorito)
}
