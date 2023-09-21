package main

import (
	interfaces "curso_golang_platzi/src/interfaces"
	mypkg "curso_golang_platzi/src/mypackage"
	"curso_golang_platzi/src/pointers"
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

type car struct {
	brand string
	year  int
	color string
	seats int
}

// type pc struct {
// 	ram     int
// 	storage int
// 	brand   string
// }

//* así se le agregan funciones a un struct
// func (myPC pc) ping() {
// 	fmt.Println(myPC.brand, "pong")
// }

//* aquí se pasa el valor por referencia (puntero)
// func (myPC *pc) duplicateRam() {
// 	// myPC.ram = myPC.ram * 2
// 	myPC.ram *= 2
// }

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
	checkIfIsPalindrome("ana")
	checkIfIsPalindrome("Ana")
	checkIfIsPalindrome("amor a roma")
	checkIfIsPalindrome("Amor a Roma")
	checkIfIsPalindrome("anal")

	// maps (platzi)
	myMap := make(map[string]int)

	myMap["Jose"] = 20
	myMap["Pepito"] = 22
	fmt.Println(myMap)

	// recorrer un map
	for k, v := range myMap {
		fmt.Println("key->", k, "value->", v)
	}

	// acceder a un valor de un map
	myValue, ok := myMap["Jose"]
	fmt.Println(myValue, ok)
	// si no existe el valor, nos devuelve un zero value
	myValue2 := myMap["Josep"]
	fmt.Println(myValue2)
	// para verificar que existe un valor, obtenemos una segunda variable (ok) al acceder a un valor de un map, que indica si este valor efectivamente existe
	myValue3, ok := myMap["Josep"]
	fmt.Println(myValue3, "Existe??", ok)

	// como instanciar un struct
	myCar := car{brand: "Ford", year: 2018}

	fmt.Println(myCar)

	// otra manera de instanciar un struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	myBestCar := car{"Porsche", 2017, "red", 2}
	fmt.Println(myBestCar)

	// pointers
	i, j := 42, 2701

	p := &i // apunta a i

	fmt.Println("puntero p al valor i:", p)
	fmt.Println("leer i por medio del puntero p:", *p)
	*p = 69
	fmt.Println("nuevo valor de i:", i)
	fmt.Println(j)

	p = &j // apunta al valor en memoria de j

	*p = *p / 37   // divide j a través del puntero
	fmt.Println(j) // 73

	// modulos
	var myCar2 mypkg.CarPublic

	myCar2.Brand = "Ferrari"
	myCar2.Year = 2020
	fmt.Println(myCar2)

	// var myCar3 mypkg.carPrivate
	// fmt.Println(myCar3)

	mypkg.PrintMessage("Hellooooo")

	//punteros #2 -> pointer is a variable that contains the address of some other value

	// https://youtu.be/NKTfNv2T0FE?t=1088
	// https://www.digitalocean.com/community/conceptual-articles/understanding-pointers-in-go-es

	a := 69
	b := &a

	// fmt.Println(a)
	fmt.Println(b)  // 0xc00000a530
	fmt.Println(*b) // 69

	*b = 138
	fmt.Println(a)

	myPc := pointers.Pc{Ram: 32, Storage: 500, Brand: "Beelink"}

	fmt.Println(myPc)
	fmt.Println(myPc)

	myPc.Ping()

	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)

	cuadrado1 := interfaces.Cuadrado{Base: 2}
	rectangulo1 := interfaces.Rectangulo{Base: 5, Altura: 2}

	interfaces.CalcularArea(cuadrado1)
	interfaces.CalcularArea(rectangulo1)

	listaDeInterfaces := []interface{}{"hello", 33, 6.9}

	fmt.Println(listaDeInterfaces...)

}
