package abstractFactoryVideo

type IAirConditioner interface {
	getType() string
}
type ColdAirConditioner struct{}

func (ac *ColdAirConditioner) getType() string {
	return "only Cold ac"
}

type HotAirConditioner struct{}

func (ac *HotAirConditioner) getType() string {
	return "only Hot ac"
}

type BothAirConditioner struct{}

func (ac *BothAirConditioner) getType() string {
	return "both Hot and cold ac"
}
