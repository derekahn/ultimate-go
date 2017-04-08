### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/PvQKHgf9jZ)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/8CtSrnTp-1))

Solution:
```go
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
		age:   45,
	}

	// Display the field values.
	fmt.Println("Name", bill.name)
	fmt.Println("Email", bill.email)
	fmt.Println("Age", bill.age)

	// Declare a variable using an anonymous struct.
	ed := struct {
		name  string
		email string
		age   int
	}{
		name:  "Ed",
		email: "ed@ardanlabs.com",
		age:   46,
	}

	// Display the field values.
	fmt.Println("Name", ed.name)
	fmt.Println("Email", ed.email)
	fmt.Println("Age", ed.age)
}
```
