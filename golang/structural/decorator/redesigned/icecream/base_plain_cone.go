package icecream

type BasePlainCone struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *BasePlainCone) GetPreperationSteps() []string {
	step := "Base Cone"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *BasePlainCone) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 10 + oldCost
}
