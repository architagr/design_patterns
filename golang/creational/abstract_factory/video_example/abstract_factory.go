package abstractFactoryVideo

import "fmt"

func CreateCarNew(country string, fuelType string) (ICar, error) {
	factory, err := CreateRegionalFactory(country)
	if err != nil {
		return nil, err
	}
	return factory.BuildCar(fuelType)
}

func CreateRegionalFactory(country string) (ICarFactory, error) {
	if country == India {
		return CreateIndianFactory(), nil
	} else if country == USA {
		return CreateUSAFactory(), nil
	} else if country == Russia {
		return CreateRussianFactory(), nil
	}
	return nil, fmt.Errorf("invalid country %s", country)
}
