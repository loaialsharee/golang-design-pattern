package main

type Shawarma struct {
    name  string
    price float64
    stars int
}

func (g *Shawarma) setName(name string) {
    g.name = name
}

func (g *Shawarma) getName() string {
    return g.name
}

func (g *Shawarma) setPrice(price float64) {
    g.price = price
}

func (g *Shawarma) getPrice() float64 {
    return g.price
}

func (g *Shawarma) setStars(stars int) {
    g.stars = stars
}

func (g *Shawarma) getStars() int {
    return g.stars
}