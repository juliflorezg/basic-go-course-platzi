package main

import (
	"fmt"
)

// chan<- significa que este channel solo va a ser de entrada de datos
func say(text string, c chan<- string) {
	// aquí le decimos que vamos a mandar el dato text por el channel c
	c <- text
}

func main() {
	fmt.Println("Hello")
	c := make(chan string)
	go say("bye", c)

	// con esta sintaxis <-c estamos diciendo que vamos a sacar los datos del channel
	fmt.Println(<-c)
}
