package main

import "fmt"

// go run src/main.go

func main() {
	// Declaración de constantes
	const pi float64 = 3.14 // no va a cambiar en el tiempo
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variables enteras
	base := 12 // variable no creada anteriormente, se le va a asignar el valor 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values -> valor por defecto que van a tener las variables si no se les ha asignado un valor

	var a int     // 0
	var b float64 // 0
	var c string  // ''
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Ejercicio => calcular el area de un cuadrado
	const baseCuadrado = 12
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
}
