package main

import (
	factorypattern "design_pattern_golang/creational/factory_pattern"
	"fmt"
	"log"
)

type PhoneType string

const (
	Pro  PhoneType = "PRO"
	Plus PhoneType = "PLUS"
)

type StorageOptions int

const (
	Gb128 StorageOptions = 128
	Gb256 StorageOptions = 256
	Gb512 StorageOptions = 512
	TB1   StorageOptions = 1024
)

type IPhone interface {
	getPhoneType() PhoneType
	getStorage() StorageOptions
	getScreenSize() int
	getCameraPixels() int
}

type Phone struct {
	screenSize, cameraPixels int
	phoneType                PhoneType
	storage                  StorageOptions
}

func (phone *Phone) getPhoneType() PhoneType {
	return phone.phoneType
}
func (phone *Phone) getStorage() StorageOptions {
	return phone.storage
}
func (phone *Phone) getScreenSize() int {
	return phone.screenSize
}
func (phone *Phone) getCameraPixels() int {
	return phone.cameraPixels
}

type IPhonePro struct {
	Phone
}

func InitIPhonePro(storage StorageOptions) IPhone {
	return &IPhonePro{
		Phone: Phone{
			screenSize:   6,
			cameraPixels: 48,
			phoneType:    Pro,
			storage:      storage,
		},
	}
}

type IPhonePlus struct {
	Phone
}

func InitIPhonePlus(storage StorageOptions) IPhone {
	return &IPhonePro{
		Phone: Phone{
			screenSize:   5,
			cameraPixels: 12,
			phoneType:    Plus,
			storage:      storage,
		},
	}
}

func main() {
	pro128 := InitIPhonePro(Gb128)
	fmt.Println(pro128.getCameraPixels())
	fmt.Println(pro128.getStorage())
	fmt.Println(pro128.getScreenSize())
	fmt.Println(pro128.getPhoneType())

	store := factorypattern.IPhoneStore{}
	phone, err := store.Order(factorypattern.India, factorypattern.Plus, factorypattern.Gb128)
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
