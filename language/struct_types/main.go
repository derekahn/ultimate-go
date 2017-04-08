// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   uint8
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{"Jack", "jack@email.env", 34}

	// Display the field values.
	fmt.Printf("{\n\tname: %+v,\n\temail: %+v,\n\tage: %+v,\n}\n", u.name, u.email, u.age)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   uint8
	}{"jill", "jill@email.env", 21}

	// Display the field values.
	fmt.Printf("{\n\tname: %+v,\n\temail: %+v,\n\tage: %+v,\n}", u2.name, u2.email, u2.age)
}
