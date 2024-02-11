package main

import "fmt"

// Stringers - customizing structs output in console!
func main() {
	myPC := pc{ram: 32, disk: 1024, brand: "Lenovo"}

	// Before defining the stringer, the output in console is {32 1024 Lenovo}
	// After defining the stringer, the output in console is { RAM: 32 GB, DISK: 1024 GB, BRAND: Lenovo }
	fmt.Println(myPC)
}

type pc struct {
	ram   int
	disk  int
	brand string
}

// The name String is neccesary to define a stringer!
func (myPC pc) String() string {
	return fmt.Sprintf("{ RAM: %d GB, DISK: %d GB, BRAND: %s }", myPC.ram, myPC.disk, myPC.brand)
}
