package ioc

import (
	"dzaba/go-dzaba/collections"
	"reflect"
)

type Closeable interface {
	Close() error
}

type ServiceScope interface {
	Closeable

	Resolve(serviceType reflect.Type) (any, error)
	ResolveAll(serviceType reflect.Type) ([]any, error)
}

type ServiceProvider interface {
	ServiceScope

	CreateScope() (ServiceScope, error)
}

type serviceProviderImpl struct {
}

func newServiceProvider() (ServiceProvider, error) {
	return &serviceProviderImpl{}, nil
}

func Resolve[T any](provider ServiceScope) (T, error) {
	var empty T
	service, err := provider.Resolve(reflect.TypeOf(empty))
	if err != nil {
		return empty, err
	}

	return service.(T), nil
}

func ResolveAll[T any](provider ServiceScope) ([]T, error) {
	var empty T
	services, err := provider.ResolveAll(reflect.TypeOf(empty))
	if err != nil {
		return nil, err
	}

	result := collections.SelectMust(services, func(element any) T {
		return element.(T)
	})

	return result, nil
}
