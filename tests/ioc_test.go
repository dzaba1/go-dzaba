package ioc

import (
	"dzaba/go-dzaba/ioc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddTransientSelf_WhenServiceIsRegisteredAsSelfTransientWithoutDependencies_ThenNewInstances(t *testing.T) {
	services := ioc.NewServiceCollection()
	err := ioc.AddTransientSelf[*firstTestInterfaceImpl](services, NewFirstTestInterface)
	assert.Nil(t, err)

	provider, err := services.BuildServiceProvder()
	assert.Nil(t, err)
	defer provider.Close()

	service1, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	service2, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	assert.NotNil(t, service1)
	assert.NotNil(t, service2)
	assert.NotEqual(t, service1.GetId(), service2.GetId())
}

func Test_AddSingletonSelf_WhenServiceIsRegisteredAsSelfSingletonWithoutDependencies_ThenTheSameInstance(t *testing.T) {
	services := ioc.NewServiceCollection()
	err := ioc.AddSingletonSelf[*firstTestInterfaceImpl](services, NewFirstTestInterface)
	assert.Nil(t, err)

	provider, err := services.BuildServiceProvder()
	assert.Nil(t, err)
	defer provider.Close()

	service1, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	service2, err := ioc.Resolve[*firstTestInterfaceImpl](provider)
	assert.Nil(t, err)

	assert.NotNil(t, service1)
	assert.NotNil(t, service2)
	assert.Equal(t, service1.GetId(), service2.GetId())
}

func Test_AddSingletonSelf_WhenServiceIsRegisteredAsSelfTransientWithSingletonDependencies_ThenTheSameInstance(t *testing.T) {
	services := ioc.NewServiceCollection()
	err := ioc.AddSingletonSelf[FirstTestInterface](services, NewFirstTestInterface)
	assert.Nil(t, err)

	err = ioc.AddTransientSelf[*secondTestInterfaceImpl](services, NewSecondTestInterface)
	assert.Nil(t, err)

	provider, err := services.BuildServiceProvder()
	assert.Nil(t, err)
	defer provider.Close()

	service1, err := ioc.Resolve[*secondTestInterfaceImpl](provider)
	assert.Nil(t, err)

	service2, err := ioc.Resolve[*secondTestInterfaceImpl](provider)
	assert.Nil(t, err)

	assert.NotNil(t, service1)
	assert.NotNil(t, service2)
	assert.Equal(t, service1.DependencyId(), service2.DependencyId())
	assert.NotEqual(t, service1.CurrentId(), service2.CurrentId())
}
