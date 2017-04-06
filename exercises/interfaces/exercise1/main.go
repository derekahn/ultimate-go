// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.
import "fmt"

// Declare the speaker interface with a single method called speak.
type speaker interface {
	speak()
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (e english) speak() {
	fmt.Println("Hello World!")
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a pointer receiver. "你好世界"
func (s *chinese) speak() {
	fmt.Println("你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker) {
	// Call the speak method from the speaker parameter.
	s.speak()
}

func main() {

	// Declare a variable of the interface speaker type
	// set to its zero value.
	var s speaker

	// Declare a variable of type english.
	e := english{}

	// Assign the english value to the speaker variable.
	s = e

	// Call the speak method against the speaker variable.
	s.speak()

	fmt.Println("\n=======\n")

	// Declare a variable of type chinese.
	c := new(chinese)

	// Assign the chinese pointer to the speaker variable.
	s = c

	// Call the speak method against the speaker variable.
	s.speak()

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	fmt.Println("\n=======\n")
	sayHello(e)
	sayHello(&e)
	sayHello(c)
}
