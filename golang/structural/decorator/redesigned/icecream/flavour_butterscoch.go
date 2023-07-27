package icecream

type FlavourButterscoch struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *FlavourButterscoch) GetPreperationSteps() []string {
	step := "1 scoop Butterscoch"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *FlavourButterscoch) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 5 + oldCost
}
