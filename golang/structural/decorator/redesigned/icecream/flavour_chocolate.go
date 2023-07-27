package icecream

type FlavourChocolate struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *FlavourChocolate) GetPreperationSteps() []string {
	step := "1 scoop Chocolate"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *FlavourChocolate) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 7 + oldCost
}
