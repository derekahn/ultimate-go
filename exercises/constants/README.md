### Exercise 1

**Part A:** Declare an untyped and typed constant and display their values.

**Part B:** Multiply two literal constants into a typed variable and display the value.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/QMrOCsHjcC)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/aUZ-7VPb9H))

Solution:
```go
package main

import "fmt"

const (
	// server is the IP address for connecting.
	server = "124.53.24.123"

	// port is the port to make that connection.
	port int16 = 9000
)

func main() {

	// Display the server information.
	fmt.Println(server)
	fmt.Println(port)

	// Calculate the number of minutes in 5320 seconds.
	minutes := 5320 / 60.0
	fmt.Println(minutes)
}
```
