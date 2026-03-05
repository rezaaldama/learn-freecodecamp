package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ---------------------------------------------------------------------------------------------------------------------
	// Array
	// ---------------------------------------------------------------------------------------------------------------------

	// Arrays are fixed-size groups of variables of the same type
	var x [6]int    // one-dimensional array
	var a [2][3]int // multi-dimensional array

	arraySixPrimes1 := [6]int{2, 3, 5, 7, 11, 13}   // array literal with len
	arraySixPrimes2 := [...]int{2, 3, 5, 7, 11, 13} // array literal without array len

	// Change data of the specified index: array[index] = value
	x[0] = 2
	x[1] = 3
	x[2] = 5
	x[3] = 7
	x[4] = 11
	x[5] = 13

	// Access data of the specified index: array[index]
	fmt.Println(x[0]) // 2

	// Return length of the array: len(array)
	fmt.Println(len(x)) // 6

	// ---------------------------------------------------------------------------------------------------------------------
	// Slice
	// ---------------------------------------------------------------------------------------------------------------------

	// Slices always have an underlying array, though it is not always specified explicitly
	var z []int // implicit slice - 0 len and 0 cap - with no values (nil)
	y := x[1:4] // explicit slice - a slice explicitly points to an array[firstIndex:lastIndex]
	// y are a slice of x with pointers = 1, len = 3, cap = 5 (cap = len(array) - pointers)

	// In addition to a pointer to an array, slice also consists of length and capacity, where length <= capacity
	fmt.Println(len(y)) // len() return length of slice
	fmt.Println(cap(y)) // cap() return cap of slice

	// Func make is used to create a slice
	slicesIntsCap10 := make([]int, 5, 10)       // implicit slice - 5 len and 10 cap - with zero values
	slicesIntsCap5 := make([]int, 5)            // implicit slice - 5 len and cap equals to len - with zero values
	slicesString := []string{"I", "Love", "Go"} // implicit slice literals - 3 len and 3 cap - with specified values
	fmt.Println(slicesIntsCap5, slicesIntsCap10, slicesString)

	// Read and write slices, just like with arrays, using bracket index
	fmt.Println(y[0])
	y[0] = 2

	// Slices are passed by reference rather than by value
	// So, any changes in its elements will be visible to the caller*
	for i := range y {
		y[i] *= 2
	}
	fmt.Println(y) // 6, 10, 14

	// Func append (variadic) is used to dynamically add elements to a slice
	// If the append() exceeds the existing array capacity, a new array will be created*
	y = append(y, 2, 3, 4)
	y = append(y, z...)

	// You can slice a slice but any unused capacity in the original slice is also shared with any subslices.
	// So, append() can creates some very odd scenarios: overwriting each other’s data.

	// Slice can hold other slices, effectively creating a matrix or a 2D slice
	rows := 4
	cols := 4
	matrix := make([][]string, 0)
	for i := 0; i < rows; i++ {
		row := make([]string, 0)
		for j := 0; j < cols; j++ {
			row = append(row, "*")
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)

	// ---------------------------------------------------------------------------------------------------------------------
	// String
	// ---------------------------------------------------------------------------------------------------------------------

	// Just an array or a slice, you can extract a single value from a string by using an index expression:
	var s string = "Hello there"
	var b byte = s[6]

	// The slice expression notation that we used with arrays and slices also works with strings:
	// Since strings are immutable, they do not have the modification problems that slices of slices do.
	var s2 string = s[4:7]

	// A single rune or byte can beconverted to a string:
	var r rune = 'x'
	var sr string = string(r)
	var by byte = 'y'
	var sby string = string(by)

	// A common bug for new Go developers is to try to make an int into a string by using a type conversion:
	var xInt int = 65
	var xStr = string(xInt) // return "65" is intended
	fmt.Println(xStr)       // return "A" instead

	xStr = fmt.Sprint(xInt)
	fmt.Println(xStr)               // return "65" with Sprint(), or
	fmt.Println(strconv.Itoa(xInt)) // return "65" with strconv library

	// ---------------------------------------------------------------------------------------------------------------------
	// Map
	// ---------------------------------------------------------------------------------------------------------------------

	// Maps are a data structure that provides key(unique)->value mapping
	// Like slices, maps are also passed by reference into function
	// Unlike arrays or slices, maps does not have capacity and orders
	var nilMap map[string]int
	nilMap2 := map[string]int{}
	ages := make(map[string]int)

	// Map literals with specified values
	ages = map[string]int{
		"John": 37,
		"Jane": 24,
	}

	fmt.Println(len(ages)) // return length of the map

	// Insert element: map[key] = value
	ages["John"] = 37
	ages["Jane"] = 24
	ages["Jane"] = 21 // overwrites 24

	// Get element: map[key]
	fmt.Println(ages["John"]) // return 37
	fmt.Println(ages["Jane"]) // return 21

	// Delete element: delete(map, key)
	fmt.Println(ages) // John and Jane
	delete(ages, "John")
	fmt.Println(ages) // Jane

	// Check key's existence (comma, ok idiom): element, boolean := map[key]
	v, ok := ages["Jane"]
	fmt.Println(v, ok) // 21, True
	v, ok = ages["John"]
	fmt.Println(v, ok) // 0, False

	// Fetching map value with a key that does not exist will return zero
	attended := map[string]bool{
		"Ann": true,
	}

	person := "Albert"
	if !attended[person] {
		fmt.Println(person, "was not attended")
	}

	// Using map as a set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// ---------------------------------------------------------------------------------------------------------------------
	// Struct
	// ---------------------------------------------------------------------------------------------------------------------

	// Maps are a convenient way to store some kinds of data, but they have limitations.
	// They don’t define an API since there’s no way to constrain a map to only allow certain keys.
	// Also, all of the values in a map must be of the same type.
	// For these reasons, maps are not an ideal way to pass data from function to function.
	// When you haverelated data that you want to group together, you should define a struct.

	// Print all unused variables
	fmt.Println(x, a, arraySixPrimes1, arraySixPrimes2)
	fmt.Println(y, z)
	fmt.Println(b, s2, sr, sby)
	fmt.Println(nilMap, nilMap2)
}
