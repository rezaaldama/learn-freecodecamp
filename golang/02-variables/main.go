package main

import "fmt"

func main() {

	// ----------------------------------------------------------------------
	// Basic variables													| 	Description							|
	// ----------------------------------------------------------------------
	// bool																			| 	true/false							|
	// string																		| 	text										|
	// int 	int8 	int16 	int32 	int64					| 	whole numbers						|
	// uint uint8 uint32 	uint64 				  			| 	positive whole numbers	|
	// uintptr																	| 	pointer arithmetic			|
	// float32 		float64												| 	decimal numbers					|
	// complex64 	complex128										| 	imaginary numbers				|
	// ----------------------------------------------------------------------

	// ------------------------------------------
	// Type			| 	Min/MaxValue								|
	// ------------------------------------------
	// int8			|		-/+ 128											|
	// int16		| 	-/+ 32768										|
	// int32  	|		-/+ 2^31 (~2 billion)				|
	// int64		|		-/+ 2^63										|
	// float32	|		-1.18*10^38 	+3.4*10^38		|
	// float64	|		-2.23*10^308	+1.8*10^308		|
	// ------------------------------------------

	// --------------------------------------------------------------------------------------------------------------
	// Type			|		Alias												|		Description																											|
	// --------------------------------------------------------------------------------------------------------------
	// Byte			|		uint8												|																																		|
	// Rune			|		int32												|																																		|
	// int			|		int32/int64  								|		unicode code point or a char; 32/64 depends on OS architecture	|
	// uint			|		uint32/int64								|		32/64 depends on OS architecture																|
	// --------------------------------------------------------------------------------------------------------------

	// Best practice - unless you have a good reason to, stick to: bool, string, int, uint, byte, rune, float64, complex128
	// Best practice - a float cannot represent a decimal value exactly, so do not use them to represent money or similar value

	// Integer operators
	// arithmetic						+ - * / %
	// assignment						+= -= *= /= %=
	// unary								++ -- - + !
	// comparison						== != > >= < <=
	// bitwise 							<< >> & | ^ &^
	// logical							&& || !

	// String operators and built-in functions
	// comparison						== != > >= < <=
	// ordering							> >= < <=
	// concenate						+
	// len(string)					count chars in string
	// string[i]						access char in string

	// Variables declaration
	// Go compiler will assign a default zero value to any variable without assigned value
	// Go compiler will detect unused variables, so we need to use blank identifier for an unused variable

	// Single declaration
	var smsSendingLimit int

	// Multiple declaration
	var (
		costPerSMS    float64
		hasPermission bool
	)

	smsSendingLimit = 10
	costPerSMS = 0.4
	hasPermission = true

	// Variable initialization with type inferrence
	var username = "admin"

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

	// Variable initialization (short)
	congrats := "Happy birthday" // inferred to be a string
	numCars := 10                // inferred to be an integer
	temperature := 0.0           // inferred to be a float
	isFunny := true              // inferred to be a bool
	complex := 0.867 + 0.5i      // inferred to be a complex
	fmt.Println(
		congrats,
		numCars,
		temperature,
		isFunny,
		complex,
	)

	// Multiple variables initialization on the same line
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	// Type conversion
	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for", accountAgeInt, "years")

	var val32 int32 = 129
	var val8 int8 = int8(val32)
	fmt.Println(val32)
	fmt.Println(val8)

	var testString = "test"
	var e byte = testString[1]
	var eString = string(e)
	fmt.Println(eString)

	// Type alias
	type FullName string
	var myName FullName = "John Doe"
	fmt.Println(myName)

	// %T - return data type of a var
	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	// Constants cannot be reassigned
	const (
		premiumPlanName = "Premium Plan"
		basicPlanName   = "Basic Plan"
	)

	fmt.Println("plan", premiumPlanName)
	fmt.Println("plan", basicPlanName)

	// A typed constant can only be directly assigned to a variable of that type
	// The following code will return compile time error
	// const untypedString = "test"; const typedInt int = untypedString;

	// Constants can be computed during compilation
	const secondInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondInMinute * minutesInHour

	fmt.Println("numbers of seconds in an hour:", secondsInHour)

	// Formatting strings
	// fmt.Printf 			Prints a formatted string to std out
	// fmt.Sprintf 			Returns a formatted string

	// %v - Interpolate the default interpretations
	fmt.Printf("I am %v years old \n", 18)
	fmt.Printf("I am %v years old \n", "way too many")

	// %s - Interpolate only a string
	fmt.Printf("I am %s years old \n", "way too many")

	// %d - Interpolate only an integer in decimal form
	fmt.Printf("I am %d years old \n", 10)

	// %f - Interpolate only a decimal form
	fmt.Printf("I am %.2f years old \n", 10.592811298461)

	const name = "Saul Goodman"
	const openRate = 30.5
	msg := fmt.Sprintf(
		"Hi %s your open rate is %.1f percent",
		name,
		openRate,
	)
	fmt.Println(msg)
}
