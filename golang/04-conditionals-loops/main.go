package main

import (
	"fmt"
	"math/rand"
)

func main() {
	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send a message length:", messageLen)

	// If-statement does not use parentheses
	if messageLen > maxMessageLen {
		fmt.Println("Message sent")
	} else if messageLen == maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message also not sent")
	}

	// If-statement with guard clauses
	if messageLen > maxMessageLen {
		fmt.Println("Message sent")
	}
	if messageLen == maxMessageLen {
		fmt.Println("Message not sent")
	}
	if messageLen < maxMessageLen {
		fmt.Println("Message also not sent")
	}

	// Conditional statements with initial statement
	// Initiated var will be limited to the statement scope

	// If-statement with initial-statement (syntatic sugar)
	if length := rand.Intn(10); length < 1 {
		fmt.Println("Email is invalid")
	}

	// Switch-statement with initial-statement (syntatic sugar)
	switch length := 10; length {
	case 1, 2, 3, 4:
		fmt.Println("a short length!")
	case 5:
		fmt.Println("exactly the right length")
	default:
		fmt.Println("a long length!")
		fmt.Println("a really long length!")
	}

	// Blank switch-statement
	switch length := 10; {
	case length < 5:
		fmt.Println("a short length!")
	case length == 5:
		fmt.Println("exactly the right length")
	default:
		fmt.Println("a long length!")
		fmt.Println("a really long length!")
	}

	// Loop-statement is only written with "for" keyword

	// There are four ways to write loops:
	// 1. A complete, C-style for
	// 		for init-statement; condition; post-statement { ... }
	//				init-statement is executed once at the start
	//				condition, if true, will continue the loop
	//				post-statement is executed at the end of the every loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// C-style for's condition can be omitted
	totalCost := 0.0
	thresh := 10.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			fmt.Printf("Threshold: %.2f\n", thresh)
			fmt.Printf("Maximum messages that can be sent: %d\n", i)
			break
		}
	}

	// 2. A condition-only for
	// 		for CONDITION { ... }
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// 3. An infinite for
	// 		for { ... }

	// for {
	//       fmt.Println("Hello")
	// }
	// Press Ctrl + C to interupt infinite loop

	// 4. for range
	// 		evenArray := []int{2, 4, 6, 8, 10}
	//		for index, value := range evenArray { fmt.Println(index, value) }
	// 		uniqueNamesMap := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	// 		for key, value := range uniqueNamesMap { fmt.Println(key, value) }
	// 		helloString := "hello"
	//		for char := range helloString { fmt.Println(string(char)) }

	// Inside the loop, "break" can be used to exit the whole loop
	// 									"continue" can be used to exit the current loop
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// By default, break and continue keywords apply to the for loop that directly contains them
	// To skip over the direct iterator, we can label our for loop
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// If i and j are both 1, we want to stop EVERYTHING
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of everything!")
				break OuterLoop // This breaks the OuterLoop, not just the inner one
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	fmt.Println("Done!")

OuterLoop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// If j is 1, skip this specific 'i' iteration entirely
			// and move to the next 'i'
			if j == 1 {
				fmt.Println("Skipping to next i...")
				continue OuterLoop2
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

}
