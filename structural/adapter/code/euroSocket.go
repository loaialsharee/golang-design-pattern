package main

import "fmt"

type EuroSocket struct{}

func (e *EuroSocket) InsertSquarePlug() {
	fmt.Println("Square plug fits into European socket.")
}
