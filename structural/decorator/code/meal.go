package main

type Meal struct {
	flight IFlight
}

func (m *Meal) getPrice() int {
	flightPrice := m.flight.getPrice()
	return flightPrice + 20
}
