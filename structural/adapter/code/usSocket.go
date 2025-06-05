package main

import "fmt"

type USSocket struct{}

func (s *USSocket) InsertRoundPlug() {
	fmt.Println("Round plug fits into US socket.")
}
