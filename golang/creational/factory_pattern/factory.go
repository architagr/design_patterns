package factorypattern

import "fmt"

type IPhoneStore struct {
	factory IFactory
}

func (store *IPhoneStore) Order(
	phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	phone, err := store.factory.CreatePhone(phoneType, memory)
	if err != nil {
		return nil, err
	}
	phone.InstallSoftwares()
	phone.Package()
	return phone, nil
}
func InitIndiaStore() *IPhoneStore {
	return &IPhoneStore{
		factory: &IPhoneIndiaFactory{},
	}
}
func InitJapanStore() *IPhoneStore {
	return &IPhoneStore{
		factory: &IPhoneJapanFactory{},
	}
}
func InitUSAStore() *IPhoneStore {
	return &IPhoneStore{
		factory: &IPhoneUSAFactory{},
	}
}
func InitStore(location MakeIn) (*IPhoneStore, error) {
	if location == India {
		return &IPhoneStore{
			factory: &IPhoneIndiaFactory{},
		}, nil
	} else if location == Japan {
		return &IPhoneStore{
			factory: &IPhoneJapanFactory{},
		}, nil
	} else if location == USA {
		return &IPhoneStore{
			factory: &IPhoneUSAFactory{},
		}, nil
	}
	return nil, fmt.Errorf("wrong location")
}

type IFactory interface {
	CreatePhone(phoneType PhoneType,
		memory StorageOptions) (IPhone, error)
}
type IPhoneIndiaFactory struct {
	IPhoneStore
}

func (indiaFactory *IPhoneIndiaFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, true, ESim, PhysicalSim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, true, ESim, PhysicalSim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}

type IPhoneJapanFactory struct {
	IPhoneStore
}

func (japanFactory *IPhoneJapanFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, false, ESim, PhysicalSim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, false, PhysicalSim, PhysicalSim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}

type IPhoneUSAFactory struct {
	IPhoneStore
}

func (usaFactory *IPhoneUSAFactory) CreatePhone(phoneType PhoneType,
	memory StorageOptions) (IPhone, error) {
	if phoneType == Pro {
		return InitIPhonePro(memory, true, ESim, ESim), nil
	} else if phoneType == Plus {
		return InitIPhonePlus(memory, true, ESim, ESim), nil
	}
	return nil, fmt.Errorf("incorrect phone type")
}
