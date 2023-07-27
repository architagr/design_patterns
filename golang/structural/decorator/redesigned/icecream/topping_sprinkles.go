package icecream

type ToppingSprinkles struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *ToppingSprinkles) GetPreperationSteps() []string {
	step := "add 10 gm sprikles"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *ToppingSprinkles) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 15 + oldCost
}
