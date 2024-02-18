package examples

import "fmt"

func ProgramFlow() {
	iterator := 0
	// This is while equivalent in Go
	for {
		fmt.Println("[While] Iterator is now: ", iterator)
		iterator++
		if iterator == 10 {
			fmt.Println("Iterator is now 10, breaking the loop")
			break
		}
	}

	// This is for over range of numbers in Go
	for i := range 10 {
		fmt.Println("[For range over numbers] Iterator is now: ", i)
	}

	// This is for over range of list in Go
	list := []string{"a", "b", "c", "d", "e"}
	for index, value := range list {
		fmt.Printf("[For range over list] Index: %d, Value: %s\n", index, value)
	}

	// This is for over range of list in Go, but we don't need index
	for _, value := range list {
		fmt.Printf("[For range over list, no index] Value: %s\n", value)
	}

	// This is standard for loop in Go
	for i := 0; i < 10; i++ {
		fmt.Println("[Standard For] Iterator is now: ", i)
	}

	for i := range 5 {
		// If statement
		if i%2 == 0 {
			fmt.Println("[IF] Number is even: ", i)
		} else {
			fmt.Println("[IF] Number is odd: ", i)
		}

		// Switch statement
		switch i%2 == 0 {
		case true:
			fmt.Println("[Switch] Number is even: ", i)
		case false:
			fmt.Println("[Switch] Number is odd: ", i)
		default:
			fmt.Println("[Switch] What is this number?????")
		}
	}
}
