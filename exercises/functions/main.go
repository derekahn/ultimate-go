// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name string
	age  int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func newUser(name string, age int) (u *user, err error) {
	if age < 21 {
		err = fmt.Errorf("User is not old enough... %+v < 21\n", age)
		return
	}
	return &user{name, age}, nil
}

func main() {
	// Use the function to create a value of type user. Check
	// the error being returned.
	u, err := newUser("jack", 24)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value that the pointer points to.
	fmt.Printf("user = %+v\n", u)

	// Call the function again and just check the error.
	u, err := newUser("jill", 18)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
