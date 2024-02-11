package mypackage

import "fmt"

// PC Struct
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Defining a method for the struct pc
func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "pong!!")
}

// Asterisk here accesses the struct attributes' memory refs using pointers
func (myPC *Pc) DuplicateRAM() {
	myPC.Ram = myPC.Ram * 2
}
