## Slices - Arrays, Slices and Maps

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

## Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.
* Slices can grow through the use of the built-in function append.

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that appends 10 values to the slice. Iterate over the slice and display each value.

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of index one and two and display the index position and value of each element in the new slice.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/sE06PRtw7h)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/3WKISOXA-L))

Solution:
```go
package main

import "fmt"

func main() {

	// Declare a nil slice of integers.
	var numbers []int

	// Append numbers to the slice.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}

	// Display each value.
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Declare a slice of strings.
	names := []string{"Bill", "Lisa", "Jim", "Cathy", "Beth"}

	// Display each index position and name.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	// Take a slice of index 1 and 2.
	slice := names[1:3]

	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
```
