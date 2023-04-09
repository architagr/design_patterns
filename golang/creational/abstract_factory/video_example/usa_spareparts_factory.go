package abstractFactoryVideo

type USASparePartsfactory struct {
}

func (spareParts *USASparePartsfactory) createWheel() IWheel {
	return &SummerWheel{}
}
func (spareParts *USASparePartsfactory) createEngine(fuelType string) IEngine {
	var enginePower IEngine = nil
	if fuelType == PETROL {
		enginePower = &MediumPowerEngine{}
	} else if fuelType == DIESEL {
		enginePower = &HighPowerEngine{}
	}
	return enginePower
}
func (spareParts *USASparePartsfactory) createAC() IAirConditioner {
	return &ColdAirConditioner{}
}
