package main

import (
	"fmt"
	"math"
)

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

	/// operaciones aritméticas

	x := 3
	y := 23
	result := x + y
	fmt.Println("Suma:", result)
	result = x - y
	fmt.Println("Resta:", result)
	result = x * y
	fmt.Println("Multiplicació:n", result)
	var division float64 = float64(x) / float64(y)
	fmt.Println("División:", division)
	modulo := x % y
	fmt.Println("Modulo:", modulo)

	//incremental
	x++
	fmt.Println("Incremental", x)
	//decremento
	x--
	fmt.Println("Decremental", x)

	// ejercicios:

	//calcular el area de un rectangulo:
	lado1 := 23
	lado2 := 48
	areaRectangulo := lado1 * lado2

	fmt.Println("Area Rectángulo:", areaRectangulo)
	//calcular el area de un trapecio:
	base1 := 23
	base2 := 48
	altura := 3.5
	areaTrapecio := (float64(base1+base2) * altura) / 2

	fmt.Println("Area Trapecio:", areaTrapecio)

	//calcular el area de un circulo:
	radio := 2.8

	areaCirculo := math.Pi * math.Pow(radio, 2) // radio al cuadrado o radio a la segunda potencia
	println("Area Circulo:", areaCirculo)
}
