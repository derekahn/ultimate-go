## Goroutines

Goroutines are functions that are created and scheduled to be run independently by the Go scheduler. The Go scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independently.
* We must always maintain an account of running goroutines and shutdown cleanly.
* Concurrency is not parallelism.
	* Concurrency is about dealing with lots of things at once.
	* Parallelism is about doing lots of things at once.

## Design Guidelines

### Concurrent Software Design

Concurrency is about managing multiple things at once. Like one person washing the dishes while they are also cooking dinner. You're making progress on both but you're only ever doing one of those things at the same time. Parallelism is about doing multiple things at once. Like one person cooking and placing dirty dishes in the sink, while another washes the dishes. They are happening at the same time.

Both you and the runtime have a responsibility in managing the concurrency of the application. You are responsible for managing these three things when writing concurrent software:

**Design Philosophy:**

* The application must startup and shutdown with integrity.
  * Know how and when every goroutine you create terminates.
  * All goroutines you create should terminate before main returns.
  * Applications should be capable of shutting down on demand, even under load, in a controlled way.
    * You want to stop accepting new requests and finish the requests you have (load shedding).
* Identify and monitor critical points of back pressure that can exist inside your application.
  * Channels, mutexes and atomic functions can create back pressure when goroutines are required to wait.
  * A little back pressure is good, it means there is a good balance of concerns.
  * A lot of back pressure is bad, it means things are imbalanced.
  * Back pressure that is imbalanced will cause:
    * Failures inside the software and across the entire platform.
    * Your application to collapse, implode or freeze.
  * Measuring back pressure is a way to measure the health of the application.
* Rate limit to prevent overwhelming back pressure inside your application.
  * Every system has a breaking point, you must know what it is for your application.
  * Applications should reject new requests as early as possible once they are overloaded.
    * Don’t take in more work than you can reasonably work on at a time.
    * Push back when you are at critical mass. Create your own external back pressure.
  * Use an external system for rate limiting when it is reasonable and practical.
* Use timeouts to release the back pressure inside your application.
  * No request or task is allowed to take forever.
  * Identify how long users are willing to wait.
  * Higher-level calls should tell lower-level calls how long they have to run.
  * At the top level, the user should decide how long they are willing to wait.
  * Use the `Context` package.
    * Functions that users wait for should take a `Context`.
      * These functions should select on `<-ctx.Done()` when they would otherwise block indefinitely.
    * Set a timeout on a `Context` only when you have good reason to expect that a function's execution has a real time limit.
    * Allow the upstream caller to decide when the `Context` should be canceled.
    * Cancel a `Context` whenever the user abandons or explicitly aborts a call.
* Architect applications to:
  * Identify problems when they are happening.
  * Stop the bleeding.
  * Return the system back to a normal state.

### Exercise 1

**Part A** Create a program that declares two anonymous functions. One that counts down from 100 to 0 and one that counts up from 0 to 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/kjtlMXkAAv)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/pUV-FPd3CE))

Solution:
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count down from 100 to 0.
		for count := 100; count >= 0; count-- {
			fmt.Printf("[A:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Count up from 0 to 100.
		for count := 0; count <= 100; count++ {
			fmt.Printf("[B:%d]\n", count)
		}

		// Tell main we are done.
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	// Display "Terminating Program".
	fmt.Println("\nTerminating Program")
}
```
