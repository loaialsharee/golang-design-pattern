package main

import "fmt"

type ISportsFactory interface {
	makeBike() IBike
	makeCar() ICar
}

func GetBrandsFactory(brand string) (ISportsFactory, error) {
	if brand == "nissan" {
		return &Nissan{}, nil
	}

	if brand == "tesla" {
		return &Tesla{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
