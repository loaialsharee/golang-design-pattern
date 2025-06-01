package main

type Tesla struct {
}

func (a *Tesla) makeCar() ICar {
	return &TeslaCar{
		Car: Car{
			logo:  "tesla",
			price: 200000,
		},
	}
}

func (a *Tesla) makeBike() IBike {
	return &TeslaBike{
		Bike: Bike{
			logo:  "tesla",
			price: 90000,
		},
	}
}
