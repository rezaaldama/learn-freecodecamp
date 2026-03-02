package main

import (
	"fmt"
	"math/rand"
)

func main() {
	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send a message length:", messageLen)

	// Go's if-statement does not use parentheses
	if messageLen > maxMessageLen {
		fmt.Println("Message sent")
	} else if messageLen == maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message also not sent")
	}

	// Go have a syntantic sugar for if-statement with initial-statement:
	// If INITIAL_STATEMENT; CONDITION {
	// ...
	// }
	// The initiated var will be limited only to the if-statement scope
	if length := rand.Intn(10); length < 1 {
		fmt.Println("Email is invalid")
	}

	// Go also have a syntatic sugar for switch-statement
	switch length := 10; length {
	case 1, 2, 3, 4:
		fmt.Println("a short length!")
	case 5:
		fmt.Println("exactly the right length")
	default:
		fmt.Println("a long length!")
	}
}
