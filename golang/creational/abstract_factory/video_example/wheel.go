package abstractFactoryVideo

type IWheel interface {
	getType() string
}
type WinterWheel struct {
}

func (wheel *WinterWheel) getType() string {
	return "Winter"
}

type SummerWheel struct {
}

func (wheel *SummerWheel) getType() string {
	return "Summer"
}
