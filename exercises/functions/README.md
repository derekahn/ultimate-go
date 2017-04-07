### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/i5wI736jpN)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/fabhfnqJ0C))

Solution:
```go
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// newUser creates and returns pointers of user type values.
func newUser() (*user, error) {
	return &user{"Bill", "bill@ardanlabs.com"}, nil
}

func main() {

	// Create a value of type user.
	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value.
	fmt.Println(*u)

	// Call the function and just check the error on the return.
	_, err = newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
