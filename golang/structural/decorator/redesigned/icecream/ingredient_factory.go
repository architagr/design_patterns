package icecream

func ingrediantFactory(ingredient IceCreamIngredient, existing IIceCreamIngrediant) IIceCreamIngrediant {
	switch ingredient {
	case BASE_PLAIN_CONE:
		return &BasePlainCone{
			ExistingIngrediant: existing,
		}
	case BASE_Chocolate_CONE:
		return &BaseChocolateCone{
			ExistingIngrediant: existing,
		}
	case FLAVOUR_BUTTERSCOTCH:
		return &FlavourButterscotch{
			ExistingIngrediant: existing,
		}
	case FLAVOUR_Chocolate:
		return &FlavourChocolate{
			ExistingIngrediant: existing,
		}
	case FLAVOUR_VANILLA:
		return &FlavourVanilla{
			ExistingIngrediant: existing,
		}
	case TOPPING_CHOCOLATE_CHIP:
		return &ToppingChocolateChip{
			ExistingIngrediant: existing,
		}
	case TOPPING_CHOCOLATE_SAUCE:
		return &ToppingChocolateSauce{
			ExistingIngrediant: existing,
		}
	case TOPPING_SPRINKLES:
		return &ToppingSprinkles{
			ExistingIngrediant: existing,
		}
	}
	return nil
}
