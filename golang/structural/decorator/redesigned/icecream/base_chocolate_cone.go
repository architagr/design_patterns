package icecream

type BaseChocolateCone struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *BaseChocolateCone) GetPreperationSteps() []string {
	step := "Chocolate Cone"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *BaseChocolateCone) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 12 + oldCost
}
