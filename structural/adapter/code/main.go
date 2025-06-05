package main

func main() {
	client := &Client{}

	usSocket := &USSocket{}
	client.PlugChargerIntoSocket(usSocket)

	euroSocket := &EuroSocket{}
	adapter := &PlugAdapter{
		euroSocket: euroSocket,
	}

	client.PlugChargerIntoSocket(adapter)
}
