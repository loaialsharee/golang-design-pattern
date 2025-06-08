package main

import "fmt"

type Shop struct {
	delivery Delivery
}

func (r *Shop) Deliver() {
	fmt.Println("Delivery request for a shop...")
	r.delivery.Deliver()
}

func (r *Shop) SetDelivery(d Delivery) {
	r.delivery = d
}
