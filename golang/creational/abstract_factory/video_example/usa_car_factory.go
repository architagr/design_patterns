package abstractFactoryVideo

type USACarFactory struct {
	sparePartfactory ISparePartsFactory
}

func CreateUSAFactory() ICarFactory {
	partfactory := new(USASparePartsfactory)
	return &USACarFactory{
		sparePartfactory: partfactory,
	}
}
func (carFactory *USACarFactory) BuildCar(fuelType string) (ICar, error) {
	err := ValidateFuelType(fuelType)
	if err != nil {
		return nil, err
	}
	carObj := new(Car)
	carObj.wheelBase = 22
	carObj.drivingSide = DrivingType_Left
	carObj.numberOfAirBags = 2
	carObj.wheel = carFactory.sparePartfactory.createWheel()
	carObj.enginePower = carFactory.sparePartfactory.createEngine(fuelType)
	carObj.ac = carFactory.sparePartfactory.createAC()
	carObj.fuelType = fuelType

	if fuelType == PETROL {
		carObj.life = 30
	} else if fuelType == DIESEL {
		carObj.life = 20
	}
	return carObj, nil
}
