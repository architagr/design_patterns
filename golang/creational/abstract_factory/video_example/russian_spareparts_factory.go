package abstractFactoryVideo

type RussianSparePartsfactory struct {
}

func (spareParts *RussianSparePartsfactory) createWheel() IWheel {
	return &WinterWheel{}
}
func (spareParts *RussianSparePartsfactory) createEngine(fuelType string) IEngine {
	var enginePower IEngine = nil
	if fuelType == PETROL {
		enginePower = &HighPowerEngine{}
	} else if fuelType == DIESEL {
		enginePower = &HighPowerEngine{}
	}
	return enginePower
}
func (spareParts *RussianSparePartsfactory) createAC() IAirConditioner {
	return &HotAirConditioner{}
}
