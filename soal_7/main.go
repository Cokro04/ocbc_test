package main

import (
	"fmt"
	"soal7/vehicle"
)

func main() {
	car := vehicle.NewCar("B 1234 CD", 15)
	car.SetMaxSpeed(180)

	truck := vehicle.NewTruck("D 9090 EF", 8, 2000)
	truck.SetMaxSpeed(130)

	car.DisplayInfo()
	fuel := car.CalculateFuelConsumption(300)
	fmt.Printf("Fuel needed: %d\n", int(fuel))

	println()

	truck.DisplayInfo()
	fuel = truck.CalculateFuelConsumption(300)
	fmt.Printf("Fuel needed: %d\n", int(fuel))
}
