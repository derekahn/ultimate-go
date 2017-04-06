// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var sliceInts []int

	// Appends numbers to the slice.
	sliceInts = append(sliceInts, 1, 2, 3)

	// Display each value in the slice.
	fmt.Println("sliceInts[0]", sliceInts[0])
	fmt.Println("sliceInts[1]", sliceInts[1])
	fmt.Println("sliceInts[2]", sliceInts[2])

	// Declare a slice of strings and populate the slice with names.
	sliceStrs := []string{"a", "b", "c"}

	fmt.Println("\n\n")

	// Display each index position and slice value.
	for i, v := range sliceStrs {
		fmt.Printf("Index = %+v Value = %s\n", i, v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	sliceTwo := sliceStrs[1:3]

	fmt.Println("\n\n")

	// Display each index position and slice values for the new slice.
	for i, v := range sliceTwo {
		fmt.Printf("Index = %+v Value = %s\n", i, v)
	}
}
