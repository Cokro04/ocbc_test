package vehicle

type Truck struct {
	BaseVehicle
	FuelEfficiency int
	CargoWeight    int
}

func NewTruck(license string, fuelEfficiency int, cargoWeight int) *Truck {
	return &Truck{
		BaseVehicle:   BaseVehicle{LicensePlate: license},
		FuelEfficiency: fuelEfficiency,
		CargoWeight:    cargoWeight,
	}
}

func (t *Truck) CalculateFuelConsumption(distance int) float64 {
	base := float64(distance) / float64(t.FuelEfficiency)
	extra := float64(t.CargoWeight) * 0.05
	return base + extra
}

func (t *Truck) SetFuelEfficiency(eff int) {
	t.FuelEfficiency = eff
}

func (t *Truck) GetFuelEfficiency() int {
	return t.FuelEfficiency
}

func (t *Truck) SetCargoWeight(weight int) {
	t.CargoWeight = weight
}

func (t *Truck) GetCargoWeight() int {
	return t.CargoWeight
}
