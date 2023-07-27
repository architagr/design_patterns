package icecream

type ToppingChocolateSauce struct {
	ExistingIngrediant IIceCreamIngrediant
}

func (ingrediant *ToppingChocolateSauce) GetPreperationSteps() []string {
	steps := []string{"dip in warm Chocolate sauce", "wait for sauce to dry"}
	if ingrediant.ExistingIngrediant != nil {
		return append(ingrediant.ExistingIngrediant.GetPreperationSteps(), steps...)
	}
	return steps
}
func (ingrediant *ToppingChocolateSauce) GetCost() int {
	oldCost := 0
	if ingrediant.ExistingIngrediant != nil {
		oldCost = ingrediant.ExistingIngrediant.GetCost()
	}
	return 6 + oldCost
}
