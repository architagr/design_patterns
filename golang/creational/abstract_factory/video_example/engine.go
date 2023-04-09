package abstractFactoryVideo

type IEngine interface {
	getPower() int
}

type HighPowerEngine struct{}

func (engine *HighPowerEngine) getPower() int {
	return 100
}

type MediumPowerEngine struct{}

func (engine *MediumPowerEngine) getPower() int {
	return 80
}

type LowPowerEngine struct{}

func (engine *LowPowerEngine) getPower() int {
	return 60
}
