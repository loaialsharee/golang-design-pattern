package main

type ICar interface {
	setLogo(logo string)
	setPrice(price int)
	getLogo() string
	getPrice() int
}

type Car struct {
	logo  string
	price int
}

func (s *Car) setLogo(logo string) {
	s.logo = logo
}

func (s *Car) getLogo() string {
	return s.logo
}

func (s *Car) setPrice(price int) {
	s.price = price
}

func (s *Car) getPrice() int {
	return s.price
}
