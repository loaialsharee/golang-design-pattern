package main

type Baggage struct {
	flight IFlight
}

func (b *Baggage) getPrice() int {
	flightPrice := b.flight.getPrice()
	return flightPrice + 145
}
