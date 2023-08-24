package pointers

import "fmt"

type Pc struct {
	Ram     int
	Storage int
	Brand   string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "pong")
}

func (myPC *Pc) DuplicateRam() {
	myPC.Ram *= 2
}
