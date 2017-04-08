## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.

## Notes

* The purpose of all programs and all parts of those programs it to transform data from one form to the other.
* Code primarily allocates, reads and writes to memory.
* Understanding type is crucial to writing good code and understanding code.
* If you don't understand the data, you don't understand the problem.
* You understand the problem better by understanding the data.
* When variables are being declared to their zero value, use the keyword var.
* When variables are being declared and initialized, use the short variable declaration operator.

### Exercise 1

**Part A:** Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/JIgjb3Ty3e)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/wNjayRMEcM))

Solution:
```go
package main

import "fmt"

func main() {

	// Declare variables that are set to their zero value.
	var age int
	var name string
	var legal bool

	// Display the value of those variables.
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// Display the value of those variables.
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
```
