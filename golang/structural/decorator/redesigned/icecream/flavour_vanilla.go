package icecream

type FlavourVanilla struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *FlavourVanilla) GetPreperationSteps() []string {
	step := "1 scoop vanilla"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *FlavourVanilla) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 3 + oldCost
}
