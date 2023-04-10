package abstractFactoryVideo

type RussianCarFactory struct {
	sparePartfactory ISparePartsFactory
}

func CreateRussianFactory() ICarFactory {
	partfactory := new(RussianSparePartsfactory)
	return &RussianCarFactory{
		sparePartfactory: partfactory,
	}
}
func (carFactory *RussianCarFactory) BuildCar(fuelType string) (ICar, error) {
	err := ValidateFuelType(fuelType)
	if err != nil {
		return nil, err
	}
	carObj := new(Car)
	carObj.wheelBase = 30
	carObj.drivingSide = DrivingType_Left
	carObj.numberOfAirBags = 4
	carObj.wheel = carFactory.sparePartfactory.createWheel()
	carObj.enginePower = carFactory.sparePartfactory.createEngine(fuelType)
	carObj.ac = carFactory.sparePartfactory.createAC()
	carObj.fuelType = fuelType

	if fuelType == PETROL {
		carObj.life = 10
	} else if fuelType == DIESEL {
		carObj.life = 7
	}
	return carObj, nil
}
