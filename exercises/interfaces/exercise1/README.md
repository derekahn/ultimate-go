### Exercise 1

**Part A** Declare an interface named speaker with a method named speak. Declare a struct named english that represents a person who speaks english and declare a struct named chinese for someone who speaks chinese. Implement the speaker interface for each struct using a value receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type speaker and assign the address of a value of type english and call the method. Do it again for a value of type chinese.

**Part B** Add a new function named sayHello that accepts a value of type speaker. Implement that function to call the speak method on the interface value. Then create new values of each type and use the function.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/adkJ3OvYpr)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/06fecJbfE4))

Solution:
```go
package main

import "fmt"

// speaker implements the voice of anyone.
type speaker interface {
	speak()
}

// english represents an english speaking person.
type english struct{}

// speak implements the speaker interface using a
// value receiver.
func (english) speak() {
	fmt.Println("Hello World")
}

// chinese represents a chinese speaking person.
type chinese struct{}

// speak implements the speaker interface using a
// pointer receiver.
func (*chinese) speak() {
	fmt.Println("你好世界")
}

func main() {

	// Declare a variable of the interface speaker type
	// set to its zero value.
	var sp speaker

	// Declare a variable of type english.
	var e english

	// Assign the english value to the speaker variable.
	sp = e

	// Call the speak method against the speaker variable.
	sp.speak()

	// Declare a variable of type chinese.
	var c chinese

	// Assign the chinese pointer to the speaker variable.
	sp = &c

	// Call the speak method against the speaker variable.
	sp.speak()

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	sayHello(english{})
	sayHello(&english{})
	sayHello(&chinese{})

	// Why does this not work?
	// sayHello(chinese{})
}

// sayHello abstracts speaking functionality.
func sayHello(sp speaker) {
	sp.speak()
}
```
