package vehicle

type Vehicle interface {
	SetMaxSpeed(speed int)
	GetMaxSpeed() int
	CalculateFuelConsumption(distance int) float64
	DisplayInfo()
}

type BaseVehicle struct {
	LicensePlate string
	MaxSpeed     int
}

func (v *BaseVehicle) SetMaxSpeed(speed int) {
	v.MaxSpeed = speed
}

func (v *BaseVehicle) GetMaxSpeed() int {
	return v.MaxSpeed
}

func (v *BaseVehicle) DisplayInfo() {
	println("License Plate:", v.LicensePlate)
	println("Max Speed:", v.MaxSpeed)
}
