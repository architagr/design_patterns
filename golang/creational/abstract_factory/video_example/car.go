package abstractFactoryVideo

type ICar interface {
	GetFuelType() string
	GetCarLife() int
	GetWheelBase() int
	GetNumberOfAirBags() int
	GetDrivingSide() string
	GetWheelType() string
	GetACType() string
	GetEnginePower() int
}

type Car struct {
	fuelType        string
	life            int
	wheelBase       int
	numberOfAirBags int
	drivingSide     string
	wheel           IWheel
	ac              IAirConditioner
	enginePower     IEngine
}

func (car *Car) GetFuelType() string {
	return car.fuelType
}

func (car *Car) GetCarLife() int {
	return car.life
}
func (car *Car) GetWheelBase() int {
	return car.wheelBase
}
func (car *Car) GetNumberOfAirBags() int {
	return car.numberOfAirBags
}
func (car *Car) GetDrivingSide() string {
	return car.drivingSide
}
func (car *Car) GetWheelType() string {
	return car.wheel.getType()
}
func (car *Car) GetACType() string {
	return car.ac.getType()
}
func (car *Car) GetEnginePower() int {
	return car.enginePower.getPower()
}

var _ ICar = new(Car)
