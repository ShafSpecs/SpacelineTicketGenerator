package main

import (
	"fmt"
	"math/rand"
)

const spacelineOne = "Space Adventures"
const spacelineTwo = "SpaceX"
const spacelineThree = "Virgin Galactic"

const marsDistance = 62_100_000
const hoursInADay = 24

func main() {
	fmt.Printf("%-16v %4v %10v %5v\n", "Spaceline", "Days", "Trip type", "Price")

	for i := 0; i < 38; i++ {
		fmt.Print("=")
	}

	fmt.Println("")

	for i := 0; i < 10; i++ {
		/* Duration Section */
		speed := rand.Intn(15) + 16

		durationInHours := (marsDistance / speed) / 3600
		totalDuration := durationInHours / 24

		/* Trip-type section */

		tripDie := rand.Intn(2)
		tripType := ""

		switch tripDie {
		case 0:
			tripType = "One-way"
		case 1:
			tripType = "Round-trip"
		}

		/* Pricing Section */

		var priceMultiplier = 1

		if tripDie == 0 {
			priceMultiplier = 1
		} else {
			priceMultiplier = 2
		}

		price := (speed + 20) * priceMultiplier

		/* Spaceline Section*/

		spaceline := ""

		switch spacelineDie := rand.Intn(3); spacelineDie {
		case 0:
			spaceline = spacelineOne
		case 1:
			spaceline = spacelineTwo
		case 2:
			spaceline = spacelineThree
		}

		/* Printing section */

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, totalDuration, tripType, price)
	}
}
