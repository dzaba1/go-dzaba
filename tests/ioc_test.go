package ioc

import (
	"dzaba/go-dzaba/ioc"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type FirstTestInterface interface {
	GetId() uuid.UUID
}

type firstTestInterfaceImpl struct {
	id uuid.UUID
}

func NewFirstTestInterface() FirstTestInterface {
	return &firstTestInterfaceImpl{
		id: uuid.New(),
	}
}

func (impl *firstTestInterfaceImpl) GetId() uuid.UUID {
	return impl.id
}

func Test_AddTransientSelf_WhenServiceIsRegisteredAsSelfTransientWithoutDependencies_ThenNewInstances(t *testing.T) {
	services := ioc.NewServiceCollection()
	err := ioc.AddTransientSelf[*firstTestInterfaceImpl](services, NewFirstTestInterface)
	assert.Nil(t, err)

	provider, err := services.BuildServiceProvder()
	assert.Nil(t, err)

	service1, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	service2, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	assert.NotNil(t, service1)
	assert.NotNil(t, service2)
	assert.Equal(t, service1.GetId(), service2.GetId())
}
