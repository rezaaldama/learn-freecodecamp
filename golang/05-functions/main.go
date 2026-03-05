package main

import "fmt"

// Function declaration has four parts:
// 1. The keyword func
// 2. The name of the function
// 3. The input parameters
// 4. The return type

// Example 1:
// Input parameters are listed in parentheses and separated by commas
// Input parameters of the same type can be declared once after the last one, assuming they are in order
func incrementSend(sendSoFar, sendToAdd int) int {
	sendSoFar += sendToAdd
	return sendSoFar
}

// Example 2:
// Variables (except pointers or pointers-like) are passed by value
// So, when a variable is passed into a function, it only receives its copy and unable to mutate the original data
func increment(x int) {
	x++
}

// Example 3:
// Variadic parameters are indicated with three dots before the type
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	// for i := 0; i < len(numbers); i++ {
	// 	total += numbers[i]
	// }

	return total
}

// Example 4:
// Multiple return values are used by wrapping them in parentheses
// If you do not need some return values, we assign it to the underscore symbol
func getNames() (string, string) {
	return "John", "Doe"
}

// Example 5:
// Return values may be given names
// Named return values area treated as new variables defined at the top of the function
func getCoords() (x, y int) {
	// Named return values are initialized with zero values

	return
	// A return statement without arguments can be used to return the named return values
	// Such naked or implicit return can harm readability
}

func main() {
	// Example 1:
	sendsSoFar := 430
	const sendToAdd = 25
	sendsSoFar = incrementSend(sendsSoFar, sendToAdd)
	fmt.Println("You have sent", sendsSoFar, "messages")

	// Example 2 - passing by value:
	x := 5
	increment(x)   // increment() creates copy of x, instead of pointing to x's data
	fmt.Println(x) // x is still 5

	// Example 3 - variadic parameters:
	fmt.Println(sumAll(1, 2, 3, 4, 5)) // a multiple integers variable

	someNumbers := []int{1, 2, 3, 4, 5} // a slice variable
	fmt.Println(sumAll(someNumbers...)) // use spread operator (triple dots) to use a slice var as parameter

	// Example 4a - multiple return values:
	firstName, lastName := getNames()
	fmt.Println("Welcome to Textio,", firstName+" "+lastName)

	// Example 4b - ignore some return values:
	myFirstName, _ := getNames()
	fmt.Println("Welcome to Textio,", myFirstName)

	//

}
