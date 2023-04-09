package main

import (
	abstractfactorypattern "design_pattern_golang/creational/abstract_factory/video_example"
	"fmt"
	"log"
)

func main() {

	carObj, err := abstractfactorypattern.CreateCar(abstractfactorypattern.India, abstractfactorypattern.PETROL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fuel type %s\n", carObj.GetFuelType())
	fmt.Printf("Power %d\n", carObj.GetEnginePower())
	fmt.Printf("Life %d\n", carObj.GetCarLife())
	fmt.Printf("Wheel base %d\n", carObj.GetWheelBase())
	fmt.Printf("AirBag # %d\n", carObj.GetNumberOfAirBags())
	fmt.Printf("Driving side %s\n", carObj.GetDrivingSide())
	fmt.Printf("Wheel type %s\n", carObj.GetWheelType())
	fmt.Printf("AC type %s\n", carObj.GetACType())

}
