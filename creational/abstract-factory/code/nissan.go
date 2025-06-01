package main

type Nissan struct {
}

func (a *Nissan) makeCar() ICar {
	return &NissanCar{
		Car: Car{
			logo:  "nissan",
			price: 140000,
		},
	}
}

func (a *Nissan) makeBike() IBike {
	return &NissanBike{
		Bike: Bike{
			logo:  "nissan",
			price: 70000,
		},
	}
}
