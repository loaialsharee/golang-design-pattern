package main

import "fmt"

type Fedex struct {
}

func (d *Fedex) Deliver() {
	fmt.Println("Delivering your parcel using FedEx Service")
}
