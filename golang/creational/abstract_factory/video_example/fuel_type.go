package abstractFactoryVideo

import "fmt"

const (
	DIESEL = "diesel"
	PETROL = "petrol"
)

func ValidateFuelType(fuelType string) error {
	if fuelType == DIESEL || fuelType == PETROL {
		return nil
	}
	return fmt.Errorf("invalid Fuel %s", fuelType)
}
