package main

import "fmt"

func main() {
	teslaFactory, _ := GetBrandsFactory("tesla")
	nissanFactory, _ := GetBrandsFactory("nissan")

	teslaCar := teslaFactory.makeCar()
	teslaBike := teslaFactory.makeBike()

	nissanCar := nissanFactory.makeCar()
	nissanBike := nissanFactory.makeBike()

	printCarDetails(teslaCar)
	printBikeDetails(teslaBike)

	printCarDetails(nissanCar)
	printBikeDetails(nissanBike)
}

func printCarDetails(t ICar) {
	fmt.Printf("Logo: %s", t.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", t.getPrice())
	fmt.Println()
}

func printBikeDetails(b IBike) {
	fmt.Printf("Logo: %s\n", b.getLogo())
	fmt.Printf("Size: %d\n", b.getPrice())
	fmt.Println("---")
}
