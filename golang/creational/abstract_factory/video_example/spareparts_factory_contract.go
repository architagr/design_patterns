package abstractFactoryVideo

type ISparePartsFactory interface {
	createWheel() IWheel
	createEngine(fuelType string) IEngine
	createAC() IAirConditioner
}
