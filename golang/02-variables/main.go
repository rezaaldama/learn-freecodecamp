package main

import "fmt"

func main() {

	// Go basic variables are:
	// bool																	=> true/false
	// string																=> text
	// int 	int8 	int16 	int32 	int64			=> whole numbers; rune is alias for int32 => unicode code point / a char in a string
	// uint uint8 uint32 	uint64 				  	=> unsigned integer or positive whole numbers; byte is alias for uint8
	// uintptr															=> pointer arithmetic
	// float32 		float64										=> decimal numbers
	// complex32 	complex64									=> imaginary numbers

	// Best practice - unless you have a good reason to, stick to following types:
	// bool, string, int, uint32, byte, rune, float64, complex128

	// Variables declaration
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

	// Shorter declaration
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
		complex)

	// Type inference
	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
	// return "The type of penniesPerText is float64"

	// Multiple initialization on the same line
	mileage, company := 80276, "Tesla"
	fmt.Println(mileage, company)

	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	// Type conversion
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")

	// Constants
	// Cannot be reassigned, but can be computed during compilation
	const premiumPlanName = "Premium Plan"
	// premiumPlanName = "Basic Plan"
	// return Compile time error: cannot assign to constant
	const basicPlanName = "Basic Plan"

	fmt.Println("plan", premiumPlanName)
	fmt.Println("plan", basicPlanName)

	const secondInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondInMinute * minutesInHour

	fmt.Println("numbers of seconds in an hour:", secondsInHour)

	// Formatting strings

	// fmt.Printf 	- Prints a formatted string to std out
	// fmt.Sprintf 	- Returns a formatted string

	// %v - Interpolate the default interpretations
	fmt.Printf("I am %v years old \n", 18)
	fmt.Printf("I am %v years old \n", "way too many")

	// %s - Interpolate a string
	fmt.Printf("I am %s years old \n", "way too many")

	// %d - Interpolate an integer in decimal form
	fmt.Printf("I am %d years old \n", 10)

	// %f - Interpolate a decimal form
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
