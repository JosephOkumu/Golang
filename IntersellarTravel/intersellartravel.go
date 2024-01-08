//This program performs calculations necessary for intersellar travel.

package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("Amount of fuel left is:", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome to planet:", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else if fuelCost > fuelRemaining {
		cantFly()
	}
	return fuelRemaining

}

func main() {
	//Test function goes here
	var fuel int
	fuel = 1000000
	var planetChoice = "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

}
