package vehicle

type Car struct {
	BaseVehicle
	FuelEfficiency int
}

func NewCar(license string, fuelEfficiency int) *Car {
	return &Car{
		BaseVehicle: BaseVehicle{LicensePlate: license},
		FuelEfficiency: fuelEfficiency,
	}
}

func (c *Car) CalculateFuelConsumption(distance int) float64 {
	return float64(distance) / float64(c.FuelEfficiency)
}

func (c *Car) SetFuelEfficiency(eff int) {
	c.FuelEfficiency = eff
}

func (c *Car) GetFuelEfficiency() int {
	return c.FuelEfficiency
}
