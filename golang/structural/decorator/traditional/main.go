package main

import (
	"decorator_traditional/icecream"
	"fmt"
)

func main() {
	icecreamObj := icecream.CreateButterscotchIcecream(1)
	fmt.Println("Steps to create the requested icecream")
	for index, step := range icecreamObj.PreperationSteps {
		fmt.Printf("Steps %d: %s\n", (index + 1), step)
	}
	fmt.Printf("Net cost $%d\n", icecreamObj.Cost)
}
