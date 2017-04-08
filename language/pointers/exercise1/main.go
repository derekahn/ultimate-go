// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

// Add imports.
import "fmt"

func main() {

	// Declare an integer variable with the value of 20.
	num := 20

	// Display the address of and value of the variable.
	fmt.Printf("Address of: %+v\n", &num)
	fmt.Printf("Value of: %+v\n", num)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	// var numPointer *int
	numPointer := &num

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Printf("\naddress: %+v\n", numPointer)
	fmt.Printf("value: %+v\n", *numPointer)
	fmt.Printf("points to: %+v\n", *numPointer)
}
