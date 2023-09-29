package main

import (
	"fmt"
)

func sendMessageToCh(text string, ch chan string) {
	ch <- text
}

func main() {
	c := make(chan string, 2)

	c <- "first message"

	// fmt.Println(len(c), cap(c)) // 1 2
	// close(c)
	c <- "last message" //! if we try to send data to a closed channel it will panic and throw an error 'panic: send on closed channel'
	// it is a good practice to close a channel once we've used it
	// close(c)
	fmt.Println(len(c), cap(c)) // 2 2

	// c <- "third message" //! error, channel full, cannot send more data to it

	//range
	for message := range c {
		fmt.Println(message)
	}

	//select
	emailCh1 := make(chan string, 1)
	emailCh2 := make(chan string, 1)

	go sendMessageToCh("message1", emailCh1)
	go sendMessageToCh("message2", emailCh2)

	for i := 0; i < 2; i++ {
		select {

		case message1 := <-emailCh1:
			fmt.Println("Este mensaje viene del channel 1", message1)
		case message2 := <-emailCh2:
			fmt.Println("Este mensaje viene del channel 2", message2)

		}
	}

}
