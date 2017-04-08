// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	prices := make(map[string]float32)

	// Initialize some data into the map.
	prices["shoes"] = 99.99
	prices["gloves"] = 68.98
	prices["sunglasses"] = 49.99

	// Display each key/value pair.
	for k, v := range prices {
		fmt.Printf("key: %+v, val: $%+v, addr: %+v\n\n", k, v, &v)
	}
}
