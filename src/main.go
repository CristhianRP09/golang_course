package main

import (
	"fmt"
	"golang_course/src/mypackage"
)

// Structs and Pointers
func main() {
	a := 50
	b := &a // & stands for pointing to memory ref of variable a

	fmt.Println("a -->", a)
	fmt.Println("b's memory ref -->", b)     // Here b shows the memory ref value (0x...)
	fmt.Println("Actual value of b -->", *b) // Asterisk (*) stands for showing the actual value of b (50)

	*b = 100 // If the value which is pointing to a memory ref is modified, the other variables pointing to that memory ref will change as well.
	fmt.Println("New value of a -->", a)

	myPC := mypackage.Pc{Ram: 16, Disk: 256, Brand: "MSI"}
	fmt.Println(myPC)
	myPC.Ping() // It calls the Ping() method in pc struct

	fmt.Println("myPC's RAM before duplication -->", myPC.Ram)
	myPC.DuplicateRAM()
	fmt.Println("myPC's RAM after duplication -->", myPC.Ram) // myPC instance has a duplicated Ram attribute!
}
