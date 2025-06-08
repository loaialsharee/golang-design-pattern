package main

type Outlet interface {
	Deliver()
	SetDelivery(Delivery)
}
