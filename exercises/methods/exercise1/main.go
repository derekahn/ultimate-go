// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
// package main

// // Add imports.
// import "fmt"

// // Declare a struct that represents a ball player.
// // Include field called name, atBats and hits.
// type player struct {
// 	name   string
// 	atBats int
// 	hits   int
// }

// // Declare a method that calculates the batting average for a player.
// func (p player) average() float32 {
// 	return float32(p.hits) / float32(p.atBats)
// }

// func main() {
// 	// Create a slice of players and populate each player
// 	// with field values.
// 	players := []player{
// 		{"Jack", 9, 3},
// 		{"Jill", 6, 2},
// 		{"Bob", 12, 4},
// 	}

// 	// Display the batting average for each player in the slice.
// 	for _, p := range players {
// 		avg := p.average()
// 		fmt.Printf("Name = %+v, average = %+v \n", p.name, avg)
// 	}
// }

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

	fmt.Println("\n\n")
	// Why did I not choose this form?
	for _, p := range ps {
		// Answer: This is a copy semantic
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}
