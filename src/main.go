package main

import "fmt"

func main() {
	// Maps
	m := make(map[string]int)
	m["X"] = 14
	m["Y"] = 20

	fmt.Println(m)

	// Iterating maps
	for i, value := range m {
		fmt.Println(i, value)
	}

	// Finding a value
	value := m["X"]
	fmt.Println(value)
	value2 := m["Z"] // Prints zero value (0 in this case) for non-existing keys
	fmt.Println(value2)
	nonExisting, ok := m["Z"] // ok says if exists or not this key (bool)
	fmt.Println(nonExisting, ok)
}
