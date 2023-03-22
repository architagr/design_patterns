package main

import (
	factorypattern "design_pattern_golang/creational/factory_pattern"
	"fmt"
	"log"
)

func main() {

	store := factorypattern.InitIndiaStore()
	phone, err := store.Order(factorypattern.Pro, factorypattern.Gb128)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("**** India factory ****")
	fmt.Println("Camera", phone.GetCameraPixels())
	fmt.Println("Storage", phone.GetStorage())
	fmt.Println("Screen Size", phone.GetScreenSize())
	fmt.Println("Phone type", phone.GetPhoneType())
	sim1, sim2 := phone.GetSimeType()
	fmt.Println("Sims", sim1, sim2)
}
