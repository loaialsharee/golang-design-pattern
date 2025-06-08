package main

import "fmt"

type Agency struct {
	delivery Delivery
}

func (r *Agency) Deliver() {
	fmt.Println("Delivery request for an agency...")
	r.delivery.Deliver()
}

func (r *Agency) SetDelivery(d Delivery) {
	r.delivery = d
}
