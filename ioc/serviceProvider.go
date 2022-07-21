package ioc

import (
	"dzaba/go-dzaba/collections"
	"dzaba/go-dzaba/utils"
	"reflect"
)

// type Closeable interface {
// 	Close() error
// }

// type ServiceScope interface {
// 	Closeable

// 	Resolve(serviceType reflect.Type) (any, error)
// 	ResolveAll(serviceType reflect.Type) ([]any, error)
// }

type ServiceProvider interface {
	Resolve(serviceType reflect.Type) (any, error)
	ResolveAll(serviceType reflect.Type) ([]any, error)

	// CreateScope() (ServiceScope, error)
}

type serviceProviderImpl struct {
}

func newServiceProvider() (ServiceProvider, error) {
	return &serviceProviderImpl{}, nil
}

func (provider *serviceProviderImpl) Resolve(serviceType reflect.Type) (any, error) {
	return nil, nil
}

func (provider *serviceProviderImpl) ResolveAll(serviceType reflect.Type) ([]any, error) {
	return nil, nil
}

func Resolve[T any](provider ServiceProvider) (T, error) {
	service, err := provider.Resolve(utils.TypeOfGeneric[T]())
	if err != nil {
		return utils.DefaultGeneric[T](), err
	}

	return service.(T), nil
}

func ResolveAll[T any](provider ServiceProvider) ([]T, error) {
	services, err := provider.ResolveAll(utils.TypeOfGeneric[T]())
	if err != nil {
		return nil, err
	}

	result := collections.SelectMust(services, func(element any) T {
		return element.(T)
	})

	return result, nil
}
