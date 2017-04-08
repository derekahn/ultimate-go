// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var i int
	var f float64
	var s string
	var b bool

	// Display the value of those variables.
	fmt.Printf("i = %+v\n", i)
	fmt.Printf("f = %+v\n", f)
	fmt.Printf("s = %+v\n", s)
	fmt.Printf("b = %+v\n", b)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	ii := 15
	ff := 3.14159265359
	ss := "I've got no strings"
	bb := true
	fmt.Println("=========================")
	// Display the value of those variables.
	fmt.Printf("ii = %+v\n", ii)
	fmt.Printf("ff = %+v\n", ff)
	fmt.Printf("ss = %+v\n", ss)
	fmt.Printf("bb = %+v\n", bb)

	// Perform a type conversion.
	pi := float32(ff)

	// Display the value of that variable.
	fmt.Println("=========================")
	fmt.Printf("pi = %+v\n", pi)
}
