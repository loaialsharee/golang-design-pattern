package main

type IShawarma interface {
    setName(name string)
    setPrice(price float64)
    setStars(stars int)
    getName() string
    getPrice() float64
    getStars() int
}