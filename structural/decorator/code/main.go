package main

import "fmt"

func main() {

	flight := &Package{}

	// Addons - Baggage Ancillaries
	flightWithBaggage := &Baggage{
		flight: flight,
	}

	// Addons - Meal Ancillaries
	flightWithBaggageAndMeal := &Meal{
		flight: flightWithBaggage,
	}

	fmt.Printf("Total price of a flight package with a baggage and meal addons is $%d\n", flightWithBaggageAndMeal.getPrice())
}
