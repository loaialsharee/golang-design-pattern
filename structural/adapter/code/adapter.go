package main

import "fmt"

type PlugAdapter struct {
	euroSocket *EuroSocket
}

func (a *PlugAdapter) InsertRoundPlug() {
	fmt.Println("Adapter converts round plug to square plug.")
	a.euroSocket.InsertSquarePlug()
}
