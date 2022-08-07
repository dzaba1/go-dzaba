package ioc

import (
	"dzaba/go-dzaba/collections"
	"dzaba/go-dzaba/utils"
	"reflect"

	"github.com/google/uuid"
)

type ServiceScope interface {
	Close() []error
	Resolve(serviceType reflect.Type) (any, error)
	ResolveAll(serviceType reflect.Type) ([]any, error)
}

type serviceScopeImpl struct {
	resolver resolver
	services map[reflect.Type][]*registrationImpl
	id       uuid.UUID
}

func newServiceScope(resolver resolver,
	services map[reflect.Type][]*registrationImpl) (ServiceScope, error) {

	return &serviceScopeImpl{
		resolver: resolver,
		services: services,
		id:       uuid.New(),
	}, nil
}

func (provider *serviceScopeImpl) Resolve(serviceType reflect.Type) (any, error) {
	return provider.resolver.resolve(serviceType, provider.id)
}

func (provider *serviceScopeImpl) ResolveAll(serviceType reflect.Type) ([]any, error) {
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

func (provider *serviceScopeImpl) Close() []error {
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
