package main

import (
	"decorator_traditional/icecream"
	"fmt"
)

func main() {
	icecreamRequest := icecream.CreateIceCreamRequest{
		Ingrediants: []icecream.IceCreamIngredient{
			icecream.BASE_PLAIN_CONE,
			icecream.FLAVOUR_BUTTERSCOCH,
			icecream.BASE_Chocolate_CONE,
			icecream.FLAVOUR_Chocolate,
			icecream.TOPPING_CHOCOLATE_SAUCE,
			icecream.TOPPING_SPRINKLES,
		},
	}
	icecreamObj := icecream.Create(icecreamRequest)
	fmt.Println("Steps to create the requested icecream")
	for index, step := range icecreamObj.PreperationSteps {
		fmt.Printf("Steps %d: %s\n", (index + 1), step)
	}
	fmt.Printf("Net cost $%d\n", icecreamObj.Cost)
}
