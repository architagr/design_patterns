package icecream

type IceCream struct {
	PreperationSteps []string
	Cost             int
}

func Create(request CreateIceCreamRequest) IceCream {
	steps := make([]string, 0, len(request.Ingrediants))
	netCost := 0
	for _, ingredient := range request.Ingrediants {
		switch ingredient {
		case BASE_PLAIN_CONE:
			steps = append(steps, "Base Cone")
			netCost += 10
		case BASE_Chocolate_CONE:
			steps = append(steps, "Chocolate Cone")
			netCost += 12
		case FLAVOUR_BUTTERSCOCH:
			steps = append(steps, "1 scoop Butterscoch")
			netCost += 5
		case FLAVOUR_Chocolate:
			steps = append(steps, "1 scoop Chocolate")
			netCost += 7
		case FLAVOUR_VANILLA:
			steps = append(steps, "1 scoop vanilla")
			netCost += 3
		case TOPPING_CHOCOLATE_CHIP:
			steps = append(steps, "add 50 gm Chocolate chip")
			netCost += 10
		case TOPPING_CHOCOLATE_SAUCE:
			steps = append(steps, "dip in warm Chocolate sauce")
			steps = append(steps, "wait for sauce to dry")
			netCost += 6
		case TOPPING_SPRINKLES:
			steps = append(steps, "add 10 gm sprikles")
			netCost += 15
		}
	}
	return IceCream{
		PreperationSteps: steps,
		Cost:             netCost,
	}
}
