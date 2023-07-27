package icecream

type ToppingChocolateChip struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *ToppingChocolateChip) GetPreperationSteps() []string {
	step := "add 50 gm Chocolate chip"
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), step)
	}
	return []string{step}
}
func (ingrediant *ToppingChocolateChip) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 10 + oldCost
}
