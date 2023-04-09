package abstractFactoryVideo

import "fmt"

func CreateCar(country string, fuelType string) (ICar, error) {
	var carObj ICar = nil
	var err error = ValidateFuelType(fuelType)
	if err != nil {
		return nil, err
	}
	if country == India {
		if fuelType == PETROL {
			carObj = &Car{
				life:            20,
				wheelBase:       22,
				numberOfAirBags: 4,
				drivingSide:     DrivingType_Right,
				wheel:           &SummerWheel{},
				ac:              &BothAirConditioner{},
				enginePower:     &LowPowerEngine{},
			}
		} else if fuelType == DIESEL {
			carObj = &Car{
				life:            15,
				wheelBase:       22,
				numberOfAirBags: 4,
				drivingSide:     DrivingType_Right,
				wheel:           &SummerWheel{},
				ac:              &BothAirConditioner{},
				enginePower:     &MediumPowerEngine{},
			}
		}
	} else if country == USA {
		if fuelType == PETROL {
			carObj = &Car{
				life:            30,
				wheelBase:       22,
				numberOfAirBags: 2,
				drivingSide:     DrivingType_Left,
				wheel:           &SummerWheel{},
				ac:              &ColdAirConditioner{},
				enginePower:     &MediumPowerEngine{},
			}
		} else if fuelType == DIESEL {
			carObj = &Car{
				life:            20,
				wheelBase:       22,
				numberOfAirBags: 2,
				drivingSide:     DrivingType_Left,
				wheel:           &SummerWheel{},
				ac:              &ColdAirConditioner{},
				enginePower:     &HighPowerEngine{},
			}
		}
	} else if country == Russia {
		if fuelType == PETROL {
			carObj = &Car{
				life:            10,
				wheelBase:       30,
				numberOfAirBags: 5,
				drivingSide:     DrivingType_Left,
				wheel:           &WinterWheel{},
				ac:              &HotAirConditioner{},
				enginePower:     &HighPowerEngine{},
			}
		} else if fuelType == DIESEL {
			carObj = &Car{
				life:            7,
				wheelBase:       30,
				numberOfAirBags: 5,
				drivingSide:     DrivingType_Left,
				wheel:           &WinterWheel{},
				ac:              &HotAirConditioner{},
				enginePower:     &HighPowerEngine{},
			}
		}
	} else {
		err = fmt.Errorf("invalid country %s", country)
	}
	return carObj, err
}
