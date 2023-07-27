package icecream

type IIceCreamIngrediant interface {
	GetPreperationSteps() []string
	GetCost() int
}

type IceCream struct {
	Ingrediant IIceCreamIngrediant
}

func Create(request CreateIceCreamRequest) IceCream {
	var ingrediantObj IIceCreamIngrediant = nil
	for _, ingredient := range request.Ingrediants {
		ingrediantObj = ingrediantFactory(ingredient, ingrediantObj)
	}
	return IceCream{
		Ingrediant: ingrediantObj,
	}
}
