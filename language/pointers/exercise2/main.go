// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import (
	"fmt"
)

// Declare a type named user.
type user struct {
	name string
	age  int
	male bool
}

// Create a function that changes the value of one of the user fields.
func changeName(u *user, name string) {
	u.name = name
}

func main() {
	// Create a variable of type user and initialize each field.
	u := user{"jack", 30, true}

	// Display the value of the variable.
	fmt.Printf("u.name %+v\nu.age %+v\nu.male %+v", u.name, u.age, u.male)

	// Share the variable with the function you declared above.
	changeName(&u, "Jill")

	// Display the value of the variable.
	fmt.Println("\n\nu.name", u.name)
}
