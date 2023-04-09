package abstractFactoryVideo

type ICarFactory interface {
	BuildCar(fuelType string) (ICar, error)
}
