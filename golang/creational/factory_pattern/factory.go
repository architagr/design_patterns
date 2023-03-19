package factorypattern

import "fmt"

type IPhoneFactory interface {
	CreatePhone(phoneType PhoneType,
		memory StorageOptions) (IPhone, error)
}

type IStore interface {
	Order(location MakeIn,
		phoneType PhoneType,
		memory StorageOptions) IPhone
}
type IPhoneStore struct {
}

func InitFactory(location MakeIn) (IPhoneFactory, error) {
	if location == India {
		return &IPhoneIndiaFactory{}, nil
	} else if location == Japan {
		return &IPhoneJapanFactory{}, nil
	} else if location == USA {
		return &IPhoneUSAFactory{}, nil
	}
	return nil, fmt.Errorf("wrong location")
}

func (store *IPhoneStore) Order(location MakeIn,
	phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	factory, err := InitFactory(location)
	if err != nil {
		return nil, err
	}
	return factory.CreatePhone(phoneType, memory)
}

type IPhoneIndiaFactory struct{}

func (indiaFactory *IPhoneIndiaFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, true, ESim, PhysicalSim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, true, ESim, PhysicalSim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}

type IPhoneJapanFactory struct{}

func (japanFactory *IPhoneJapanFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, false, ESim, PhysicalSim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, false, PhysicalSim, PhysicalSim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}

type IPhoneUSAFactory struct{}

func (usaFactory *IPhoneUSAFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, true, ESim, ESim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, true, ESim, ESim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}
