package main

import "fmt"

type Vehicle interface {
	StartEngine() error
}

type Chargeable interface {
	ChargeBattery(percent int) error
}

// ElectricCar satisfies BOTH interfaces
type ElectricCar struct{}

func (e *ElectricCar) StartEngine() error {
	fmt.Println("Electric motor started")
	return nil
}

func (e *ElectricCar) ChargeBattery(percent int) error {
	fmt.Printf("Charging to %d%%\n", percent)
	return nil
}

// GasCar satisfies only Vehicle — honest contract
type GasCar struct{}

func (g *GasCar) StartEngine() error {
	fmt.Println("Vroom!")
	return nil
}

// PrepareVehicle works for ALL vehicles, gas or electric
func PrepareVehicle(v Vehicle) error {
	return v.StartEngine()
}

// PrepareChargeable only accepts things that can actually charge
func PrepareChargeable(c Chargeable) error {
	return c.ChargeBattery(100)
}

func main() {
	// Concrete types
	electric := &ElectricCar{}
	gas := &GasCar{}

	// Any vehicle can be prepared — gas or electric
	vehicles := []Vehicle{electric, gas}
	for _, v := range vehicles {
		if err := PrepareVehicle(v); err != nil {
			fmt.Printf("failed to prepare vehicle: %v\n", err)
		}
	}

	// Only chargeables go here — GasCar excluded naturally
	// Compiler will highlight an error if you attempt to add gas type
	chargeables := []Chargeable{electric}
	for _, c := range chargeables {
		if err := PrepareChargeable(c); err != nil {
			fmt.Printf("failed to charge: %v\n", err)
		}
	}
}
