package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(text)
}

func main() {

	// agregar un wait group
	var wg sync.WaitGroup
	fmt.Println("hello")
	wg.Add(2)
	go say("Bye", &wg)

	// wg.Add(1)

	go func(text string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(text)
	}("jeje", &wg)

	wg.Wait()
	// time.Sleep(time.Second * 1)
}
