package main

import "fmt"

type Dhl struct {
}

func (d *Dhl) Deliver() {
	fmt.Println("Delivering your parcel using DHL Service")
}
