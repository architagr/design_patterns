package factorypattern

type SimType string

const (
	ESim        SimType = "E-Sim"
	PhysicalSim SimType = "PhysicalSim"
)

type MakeIn string

const (
	India MakeIn = "India"
	Japan MakeIn = "Japan"
	USA   MakeIn = "USA"
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
	GetPhoneType() PhoneType
	GetStorage() StorageOptions
	GetScreenSize() int
	GetCameraPixels() int
	GetSimeType() (sim1 SimType, sim2 SimType)
}

type Phone struct {
	screenSize, cameraPixels int
	phoneType                PhoneType
	storage                  StorageOptions
	sim1, sim2               SimType
	canTurnOffShutterSound   bool
}

func (phone *Phone) GetPhoneType() PhoneType {
	return phone.phoneType
}
func (phone *Phone) GetStorage() StorageOptions {
	return phone.storage
}
func (phone *Phone) GetScreenSize() int {
	return phone.screenSize
}
func (phone *Phone) GetCameraPixels() int {
	return phone.cameraPixels
}
func (phone *Phone) GetSimeType() (sim1 SimType, sim2 SimType) {
	sim1, sim2 = phone.sim1, phone.sim2
	return
}

type IPhonePro struct {
	Phone
}

func InitIPhonePro(storage StorageOptions, canTurnOddShutter bool, sim1, sim2 SimType) IPhone {
	return &IPhonePro{
		Phone: Phone{
			screenSize:             6,
			cameraPixels:           48,
			phoneType:              Pro,
			storage:                storage,
			canTurnOffShutterSound: canTurnOddShutter,
			sim1:                   sim1,
			sim2:                   sim2,
		},
	}
}

type IPhonePlus struct {
	Phone
}

func InitIPhonePlus(storage StorageOptions, canTurnOddShutter bool, sim1, sim2 SimType) IPhone {
	return &IPhonePro{
		Phone: Phone{
			screenSize:             5,
			cameraPixels:           12,
			phoneType:              Plus,
			storage:                storage,
			canTurnOffShutterSound: canTurnOddShutter,
			sim1:                   sim1,
			sim2:                   sim2,
		},
	}
}
