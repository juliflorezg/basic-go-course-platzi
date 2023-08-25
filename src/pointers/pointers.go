package pointers

import "fmt"

type Pc struct {
	Ram     int
	Storage int
	Brand   string
}

// Stringers: con esta función sobreescribimos la función String del struct, para que al imprimir el struct en pantalla, se pueda ver el resultado formateado:
//* fmt.Println(myPc) -> Tengo 32 gb de ram, 500 gb de almacenamiento y la marca es Beelink
func (myPC Pc) String() string {
	return fmt.Sprintf("Tengo %d gb de ram, %d gb de almacenamiento y la marca es %s", myPC.Ram, myPC.Storage, myPC.Brand)
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "pong")
}

func (myPC *Pc) DuplicateRam() {
	myPC.Ram *= 2
}
