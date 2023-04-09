package abstractFactoryVideo

type IndianCarFactory struct {
	sparePartfactory ISparePartsFactory
}

func CreateIndianFactory() ICarFactory {
	partfactory := new(IndianSparePartsfactory)
	return &IndianCarFactory{
		sparePartfactory: partfactory,
	}
}
func (carFactory *IndianCarFactory) BuildCar(fuelType string) (ICar, error) {
	err := ValidateFuelType(fuelType)
	if err != nil {
		return nil, err
	}
	carObj := new(Car)
	carObj.wheelBase = 22
	carObj.drivingSide = DrivingType_Right
	carObj.numberOfAirBags = 4
	carObj.wheel = carFactory.sparePartfactory.createAC()
	carObj.enginePower = carFactory.sparePartfactory.createEngine(fuelType)
	carObj.ac = carFactory.sparePartfactory.createAC()
	carObj.fuelType = fuelType
	if fuelType == PETROL {
		carObj.life = 20
	} else if fuelType == DIESEL {
		carObj.life = 15
	}
	return carObj, nil
}
