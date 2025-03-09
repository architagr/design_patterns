package icecream

type FlavourButterscotch struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *FlavourButterscotch) GetPreperationSteps() []string {
	step := "1 scoop Butterscotch"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *FlavourButterscotch) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 5 + oldCost
}
