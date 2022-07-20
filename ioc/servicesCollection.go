package ioc

import "reflect"

type ServicesCollection interface {
	BuildServiceProvder() (ServiceProvider, error)
	AddTransientSelf(selfType reflect.Type, ctorFunc any) error
}

type servicesCollectionImpl struct {
}

func NewServiceCollection() ServicesCollection {
	return &servicesCollectionImpl{}
}

func (services *servicesCollectionImpl) BuildServiceProvder() (ServiceProvider, error) {
	return newServiceProvider()
}

func (services *servicesCollectionImpl) AddTransientSelf(selfType reflect.Type, ctorFunc any) error {
	err := services.validateSelfTypes(selfType, ctorFunc)
	if err != nil {
		return err
	}

}

func (services *servicesCollectionImpl) validateSelfTypes(selfType reflect.Type, ctorFunc any) error {

}

func AddTransientSelf[T any](services ServicesCollection, ctorFunc any) error {
	var empty T
	return services.AddTransientSelf(reflect.TypeOf(empty), ctorFunc)
}
