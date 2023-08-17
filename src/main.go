package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func checkIfIsPalindrome(text string) {
	var myNewString string

	for i := len(text) - 1; i >= 0; i-- {
		// fmt.Println(strings.ToLower(string(text[i])))
		// fmt.Println(strings.ToLower(text[i]))
		myNewString += strings.ToLower(string(text[i]))
	}

	// fmt.Println(myNewString)

	if strings.ToLower(text) == myNewString {
		fmt.Println("it IS a palindrome")
	} else {
		fmt.Println("it is NOT a palindrome")
	}
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

	truevar := true
	// falsevar := false
	// fmt.Printf("true %t\n", truevar) //~ %t para valores boooleanos b
	fmt.Printf("true %T\n", truevar)
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
	fmt.Printf("m type: %T \n", m)

	//** https://devhints.io/go

	//? condicionales

	// reto condicionales
	// funcion que verifica si un numero es par o impar
	checkOddOrEven := func(a int) string {
		if a%2 == 0 {
			return "el numero es par"
		} else {
			return "el numero es impar"
		}
	}

	// https://forum.golangbridge.org/t/writing-function-inside-and-outside-of-func-main/13268/3

	fmt.Println(checkOddOrEven(2))
	fmt.Println(checkOddOrEven(3))
	fmt.Println(checkOddOrEven(4))

	checkUserAndPassword := func(user, password string) (result bool) {
		if user == "pepito" && password == "1234jejeje" {
			return true
		} else {
			return false
		}
	}

	fmt.Println("are user & password ok:", checkUserAndPassword("lol", "fadsf"))
	fmt.Println("are user & password ok:", checkUserAndPassword("pepito", "1234jejeje"))

	// convertir un numero que viene en texto a entero
	// convertedValue, err := strconv.Atoi("dsds") //! strconv.Atoi: parsing "dsds": invalid syntax
	convertedValue, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("convertedValue", convertedValue)

	//* switch

	// modulo := 4 % 2

	// switch modulo {
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// switch sin condición

	value := 10

	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("Está entre 0 y 100")
	}

	//> keywords importantes

	//? defer
	// defer va a posponer la ejecución hasta justo antes que el programa main termine
	// Se puede usar para cuando nos conectamos a una BD, usamos defer para cerrar la conexión a la BD
	// También para cuando usamos un archivo, ponemos en defer que cierre el archivo
	defer fmt.Println("lol")
	fmt.Println("jejeje")

	//? continue

	for i := 0; i <= 10; i++ {
		if i == 2 {
			fmt.Println("Continue in 2")
			continue // continue no permite la ejecución de las siguientes lineas de código en el bucle, salta o --continúa-- a la siguiente iteración
			// en este caso, cuando i == 2, el println de 2 más abajo, no se va a ejecutar
		}

		fmt.Println(i)

		if i == 7 {
			break // break sale del ciclo actual
		}
	}

	//> arrays

	//* array de posiciones fijas, inmutable (no se puede agregar un elemento), por defecto va a tener 0 values
	var array [4]int
	fmt.Println(array) // [0 0 0 0]

	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//> slice

	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//> slicing

	fmt.Println("Slicing:::")
	// metodos en el slice
	fmt.Println(slice[0])   // posición 0 -> 0
	fmt.Println(slice[:3])  //  hasta la posición 3, no inclusivo al final -> [0 1 2]
	fmt.Println(slice[2:4]) //  entre posiciones 2 y 4, inclusivo al inicio, no inclusivo al final -> [2, 3]
	fmt.Println(slice[4:])  // a partir de la posición 4, inclusivo al inicio -> [4 5 6]

	//* como adicionar elementos a un slice
	slice = append(slice, 7)
	fmt.Println(slice)
	//* como adicionar otra lista a un slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//> recorrer slice con range
	stringSlice := []string{"hola", "que", "hace"}
	for i, value := range stringSlice {
		fmt.Println("index:", i, "value:", value)
	}

	// si solo necesito el indice
	for i := range stringSlice {
		fmt.Println("index:", i)
	}
	// si solo necesito el valor
	for _, value := range stringSlice {
		fmt.Println("value:", value)
	}

	// crear una función que verifique si un string es un palindromo
	// checkIfIsPalindrome("ana")
	checkIfIsPalindrome("Ana")
	// checkIfIsPalindrome("amor a roma")
	// checkIfIsPalindrome("anal")

}
