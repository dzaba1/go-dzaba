package ioc

import (
	"dzaba/go-dzaba/collections"
	"dzaba/go-dzaba/utils"
	"reflect"

	"github.com/google/uuid"
)

type Closeable interface {
	Close() error
}

type ServiceScope interface {
	Close() []error
	Resolve(serviceType reflect.Type) (any, error)
	ResolveAll(serviceType reflect.Type) ([]any, error)
}

type ServiceProvider interface {
	ServiceScope

	CreateScope() (ServiceScope, error)
}

type serviceProviderImpl struct {
	resolver resolver
	services map[reflect.Type][]*registrationImpl
	id       uuid.UUID
}

func newServiceProvider(resolver resolver,
	services map[reflect.Type][]*registrationImpl) (ServiceProvider, error) {

	return &serviceProviderImpl{
		resolver: resolver,
		services: services,
		id:       uuid.New(),
	}, nil
}

func (provider *serviceProviderImpl) Resolve(serviceType reflect.Type) (any, error) {
	return provider.resolver.resolve(serviceType, provider.id)
}

func (provider *serviceProviderImpl) ResolveAll(serviceType reflect.Type) ([]any, error) {
	result := []any{}

	for _, regs := range provider.services {
		for _, reg := range regs {
			service, err := provider.resolver.resolveRegistration(reg, provider.id)
			if err != nil {
				return nil, err
			}
			result = append(result, service)
		}
	}

	return result, nil
}

func (provider *serviceProviderImpl) Close() []error {
	errors := []error{}

	for _, registrations := range provider.services {
		for _, reg := range registrations {
			inst := reg.lifetimeManager.Instance(provider.id)
			if inst != nil {
				cast, ok := inst.(Closeable)
				if ok {
					err := cast.Close()
					if err != nil {
						errors = append(errors, err)
					}
				}
				reg.lifetimeManager.ClearInstance(provider.id)
			}
		}
	}

	return errors
}

func (provider *serviceProviderImpl) CreateScope() (ServiceScope, error) {
	return newServiceProvider(provider.resolver, provider.services)
}

func Resolve[T any](provider ServiceScope) (T, error) {
	service, err := provider.Resolve(utils.TypeOfGeneric[T]())
	if err != nil {
		return utils.DefaultGeneric[T](), err
	}

	return service.(T), nil
}

func ResolveAll[T any](provider ServiceScope) ([]T, error) {
	services, err := provider.ResolveAll(utils.TypeOfGeneric[T]())
	if err != nil {
		return nil, err
	}

	result := collections.SelectMust(services, func(element any) T {
		return element.(T)
	})

	return result, nil
}
