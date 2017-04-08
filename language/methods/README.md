## Methods

Methods are functions that are declared with a receiver which binds the method to a type. Methods can be used to operate on values or pointers of that type.

## Notes

* Methods are functions that declare a receiver variable.
* Receivers bind a method to a type and can use value or pointer semantics.
* Value semantics mean a copy of the value is passed across program boundaries.
* Pointer semantics mean a copy of the values address is passed across program boundaries.
* Stick to a single semantic for a given type and be consistent.

## Quotes

_"Methods are valid when it is practical or reasonable for a piece of data to expose a capability." - William Kennedy_

### Exercise 1

Declare a struct that represents a baseball player. Include name, atBats and hits. Declare a method that calculates a players batting average. The formula is Hits / AtBats. Declare a slice of this type and initialize the slice with several players. Iterate over the slice displaying the players name and batting average.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/IG5uqVRTrc)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/1vr9fCLEO8))

Solution:
```go
package main

import "fmt"

// player represents a person in the game.
type player struct {
	name   string
	atBats int
	hits   int
}

// average calculates the batting average for a player.
func (p *player) average() float64 {
	if p.atBats == 0 {
		return 0.0
	}

	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a few players.
	ps := []player{
		{"bill", 10, 7},
		{"jim", 12, 6},
		{"ed", 6, 4},
	}

	// Display the batting average for each player.
	for i := range ps {
    // Answer: This is a value/pointer semantic
		fmt.Printf("%s: AVG[.%.f]\n", ps[i].name, ps[i].average()*1000)
	}

	// Why did I not choose this form?
	for _, p := range ps {
    // Answer: This is a copy semantic
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}

```
