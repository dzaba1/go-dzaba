package ioc

import (
	"dzaba/go-dzaba/collections"

	"github.com/google/uuid"
)

type FirstTestInterface interface {
	GetId() uuid.UUID
}

type SecondTestInterface interface {
	DependencyId() uuid.UUID
	CurrentId() uuid.UUID
}

type AggregatedInterfaces interface {
	GetDependencyIds() []uuid.UUID
}

type firstTestInterfaceImpl struct {
	id uuid.UUID
}

type firstTestInterfaceSecondImpl struct {
	id uuid.UUID
}

type secondTestInterfaceImpl struct {
	dependency FirstTestInterface
	id         uuid.UUID
}

type aggregatedInterfacesImpl struct {
	dependency []FirstTestInterface
}

func NewFirstTestInterface() FirstTestInterface {
	return &firstTestInterfaceImpl{
		id: uuid.New(),
	}
}

func NewFirstTestInterfaceSecondImpl() FirstTestInterface {
	return &firstTestInterfaceSecondImpl{
		id: uuid.New(),
	}
}

func NewSecondTestInterface(dependency FirstTestInterface) SecondTestInterface {
	return &secondTestInterfaceImpl{
		id:         uuid.New(),
		dependency: dependency,
	}
}

func NewAggregatedInterfaces(dependency []FirstTestInterface) AggregatedInterfaces {
	return &aggregatedInterfacesImpl{
		dependency: dependency,
	}
}

func (impl *firstTestInterfaceImpl) GetId() uuid.UUID {
	return impl.id
}

func (impl *firstTestInterfaceSecondImpl) GetId() uuid.UUID {
	return impl.id
}

func (impl *secondTestInterfaceImpl) DependencyId() uuid.UUID {
	return impl.dependency.GetId()
}

func (impl *secondTestInterfaceImpl) CurrentId() uuid.UUID {
	return impl.id
}

func (impl *aggregatedInterfacesImpl) GetDependencyIds() []uuid.UUID {
	return collections.SelectMust(impl.dependency, func(element FirstTestInterface) uuid.UUID {
		return element.GetId()
	})
}
