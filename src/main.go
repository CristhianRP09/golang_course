package main

import "fmt"

func main() {
	module := 4 % 2

	switch module {
	case 0:
		fmt.Println("Is even!")
	default:
		fmt.Println("Is odd!")
	}

	// In a short way
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Is even!")
	default:
		fmt.Println("Is odd!")
	}

	// Switch sentence without condition
	value := 120
	switch {
	case value > 100:
		fmt.Println("Greater than 100")
	case value < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("Within range from 0 to 100")
	}

	// defer keyword
	fmt.Println("Hello ... {")
	defer fmt.Println("... World }")

	// continue and break keywords
	for i := 0; i < 10; i++ {
		if i != 2 {
			fmt.Println(i)
		} else {
			fmt.Println("It's 2")
			continue
		}

		if i == 8 {
			fmt.Println("It's 8... END")
			break
		}
	}
}
