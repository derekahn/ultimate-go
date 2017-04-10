// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("FIRST ROUTINE", ch)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("SECOND ROUTINE", ch)
		wg.Done()
	}()

	// Send a value to start the counting.
	ch <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(title string, signal chan int) {
	for {
		// Wait for the value to be sent.
		// If the channel was closed, return.
		value, ok := <-signal
		if !ok {
			fmt.Printf("run %+v", title)
			return
		}

		// Display the value.
		fmt.Printf("%s signaling %d\n\n", title, value)

		// Terminate when the value is 10.
		if value == 10 {
			fmt.Printf("CLOSED %+v\n", title)
			close(signal)
			return
		}

		// Increment the value and send it
		// over the channel.
		signal <- (value + 1)
	}
}
