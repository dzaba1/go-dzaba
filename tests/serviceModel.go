package ioc

import (
	"github.com/google/uuid"
)

type FirstTestInterface interface {
	GetId() uuid.UUID
}

type SecondTestInterface interface {
	DependencyId() uuid.UUID
	CurrentId() uuid.UUID
}

type firstTestInterfaceImpl struct {
	id uuid.UUID
}

type secondTestInterfaceImpl struct {
	dependency FirstTestInterface
	id         uuid.UUID
}

func NewFirstTestInterface() FirstTestInterface {
	return &firstTestInterfaceImpl{
		id: uuid.New(),
	}
}

func NewSecondTestInterface(dependency FirstTestInterface) SecondTestInterface {
	return &secondTestInterfaceImpl{
		id:         uuid.New(),
		dependency: dependency,
	}
}

func (impl *firstTestInterfaceImpl) GetId() uuid.UUID {
	return impl.id
}

func (impl *secondTestInterfaceImpl) DependencyId() uuid.UUID {
	return impl.dependency.GetId()
}

func (impl *secondTestInterfaceImpl) CurrentId() uuid.UUID {
	return impl.id
}
