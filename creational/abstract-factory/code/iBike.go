package main

type IBike interface {
	setLogo(logo string)
	setPrice(price int)
	getLogo() string
	getPrice() int
}

type Bike struct {
	logo  string
	price int
}

func (s *Bike) setLogo(logo string) {
	s.logo = logo
}

func (s *Bike) getLogo() string {
	return s.logo
}

func (s *Bike) setPrice(price int) {
	s.price = price
}

func (s *Bike) getPrice() int {
	return s.price
}
