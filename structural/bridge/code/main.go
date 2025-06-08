package main

import "fmt"

func main() {

	dhl := &Dhl{}
	fedex := &Fedex{}

	shop := &Shop{}

	shop.SetDelivery(dhl)
	shop.Deliver()
	fmt.Println()

	shop.SetDelivery(fedex)
	shop.Deliver()
	fmt.Println()

	agency := &Agency{}

	agency.SetDelivery(dhl)
	agency.Deliver()
	fmt.Println()

	agency.SetDelivery(fedex)
	agency.Deliver()
	fmt.Println()
}
