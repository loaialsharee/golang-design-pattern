package main

import "fmt"

func main() {
    chicken, _ := getShawarma("chicken")
    beef, _ := getShawarma("beef")

    printDetails(chicken)
    printDetails(beef)
}

func printDetails(s IShawarma) {
    fmt.Printf("Name: %s", s.getName())
    fmt.Println()
    fmt.Printf("Price: %.2f", s.getPrice())
    fmt.Println()
    fmt.Printf("Stars: %d", s.getStars())
	fmt.Println()
    fmt.Println("---")
}