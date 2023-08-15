package main

import (
	"fmt"
)

func printMessage(message string) {
	fmt.Println(message)
}

// func tripleArgument(a int, b int, c string) {
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func doubledValue(a int) int {
	return a * 2
}

func returnsTwoValues(a int) (b, c int) {
	return a, a * 2
}

// go run src/main.go

func main() {
	// Declaración de constantes
	const pi float64 = 3.14 // no va a cambiar en el tiempo
	const pi2 = 3.1415

	// fmt.Println("pi", pi)
	// fmt.Println("pi2", pi2)

	// //Declaración de variables enteras
	// base := 12 // variable no creada anteriormente, se le va a asignar el valor 12
	// var height int = 14
	// var area int

	// fmt.Println(base, height, area) // area va a ser 0 por el valor por defecto

	// // Zero values -> valor por defecto que van a tener las variables si no se les ha asignado un valor

	// var a int     // 0
	// var b float64 // 0
	// var c string  // ''
	// var d bool    // false

	// fmt.Println(a, b, c, d)

	// // Ejercicio => calcular el area de un cuadrado
	// const baseCuadrado = 12
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println(areaCuadrado)

	// /// operaciones aritméticas

	// x := 3
	// y := 23
	// result := x + y
	// fmt.Println("Suma:", result)
	// result = x - y
	// fmt.Println("Resta:", result)
	// result = x * y
	// fmt.Println("Multiplicació:n", result)
	// var division float64 = float64(x) / float64(y)
	// fmt.Println("División:", division)
	// modulo := x % y
	// fmt.Println("Modulo:", modulo)

	// //incremental
	// x++
	// fmt.Println("Incremental", x)
	// //decremento
	// x--
	// fmt.Println("Decremental", x)

	// // ejercicios:

	// //calcular el area de un rectangulo:
	// lado1 := 23
	// lado2 := 48
	// areaRectangulo := lado1 * lado2

	// fmt.Println("Area Rectángulo:", areaRectangulo)
	// //calcular el area de un trapecio:
	// base1 := 23
	// base2 := 48
	// altura := 3.5
	// areaTrapecio := (float64(base1+base2) * altura) / 2

	// fmt.Println("Area Trapecio:", areaTrapecio)

	// //calcular el area de un circulo:
	// radio := 2.8

	// areaCirculo := math.Pi * math.Pow(radio, 2) // radio al cuadrado o radio a la segunda potencia
	// println("Area Circulo:", areaCirculo)

	// // Paquete fmt

	// helloMessage := "Hello"
	// worldMessage := "World"

	// fmt.Println(helloMessage, worldMessage)
	// fmt.Println(helloMessage, worldMessage)

	// // Printf
	// name := "platzi"
	// courses := 500
	// fmt.Printf("%s tiene más de %d cursos\n", name, courses) //%s se usa para decir que ahí va a ir un string y %d es para dígitos
	// fmt.Printf("%v tiene más de %v cursos\n", name, courses) //%v se usa para decir que no se sabe el tipo de dato de antemano

	// // Sprintf
	// message := fmt.Sprintf("%s tiene más de %d cursos`", name, courses) // Sprintf genera una string, pero no la imprime en consola
	// fmt.Println("message:", message)

	// // Para conocer el tipo de datos
	// fmt.Printf("helloMessage: %T\n", helloMessage)
	// fmt.Printf("courses: %T\n", courses)

	// truevar := true
	// falsevar := false
	// fmt.Printf("true %t\n", truevar) //~ %t para valores boooleanos b
	// fmt.Printf("false %t\n", falsevar)

	// //? funciones
	// printMessage("Hello world")
	// tripleArgument(1, 2, "Hi!!!")

	// value := doubledValue(34)

	// fmt.Println(value)

	// value1, value2 := returnsTwoValues(2)

	// fmt.Println("value1 y value2:", value1, value2)

	// // si quiero ignorar un valor de retorno:
	// _, value3 := returnsTwoValues(5)
	// fmt.Println("value3:", value3)

	//> Flujo de control y condicionales

	// for condicional

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println("counter:", counter)
		counter++
	}

	// for forever

	// counterForever := 0

	// for {
	// 	fmt.Println("infinito", counterForever)
	// 	counterForever++
	// }

	// ciclo for inverso
	for j := 15; j > 0; j-- {
		fmt.Println("decremento de j:", j)
	}

	// for range -> para colección en un objeto
	listaDePares := []int{2, 4, 6, 8, 10}

	fmt.Println(listaDePares)

	for i, par := range listaDePares {
		fmt.Printf("Posición: %d, numero par: %d \n", i, par)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

	// 0 103 -> index & rune
	// 1 111
	// 2 108
	// 3 97
	// 4 110
	// 5 103

	// maps
	var m map[string]int
	m = make(map[string]int, 2)
	fmt.Println(m)

	m["uno"] = 1
	m["dos"] = 2
	m["three"] = 3

	fmt.Println(m)
	fmt.Printf("m type: %T", m)

	//** https://devhints.io/go

}
