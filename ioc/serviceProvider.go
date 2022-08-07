package ioc

import (
	"reflect"

	"github.com/google/uuid"
)

type Closeable interface {
	Close() error
}

type ServiceProvider interface {
	ServiceScope

	CreateScope() (ServiceScope, error)
}

type serviceProviderImpl struct {
	serviceScopeImpl
}

func newServiceProvider(resolver resolver,
	services map[reflect.Type][]*registrationImpl) (ServiceProvider, error) {

	p := &serviceProviderImpl{}
	p.resolver = resolver
	p.services = services
	p.id = uuid.New()

	return p, nil
}

func (provider *serviceProviderImpl) CreateScope() (ServiceScope, error) {
	return newServiceScope(provider.resolver, provider.services)
}
