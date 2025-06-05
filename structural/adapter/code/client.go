package main

import "fmt"

type Client struct{}

func (c *Client) PlugChargerIntoSocket(socket Socket) {
	fmt.Println("Client plugs round charger into socket.")
	socket.InsertRoundPlug()
}
