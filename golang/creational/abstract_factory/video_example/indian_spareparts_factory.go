package abstractFactoryVideo

type IndianSparePartsfactory struct {
}

func (spareParts *IndianSparePartsfactory) createWheel() IWheel {
	return &SummerWheel{}
}
func (spareParts *IndianSparePartsfactory) createEngine(fuelType string) IEngine {
	var enginePower IEngine = nil
	if fuelType == PETROL {
		enginePower = &LowPowerEngine{}
	} else if fuelType == DIESEL {
		enginePower = &MediumPowerEngine{}
	}
	return enginePower
}
func (spareParts *IndianSparePartsfactory) createAC() IAirConditioner {
	return &BothAirConditioner{}
}
