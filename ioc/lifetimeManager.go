package ioc

import "github.com/google/uuid"

type LifetimeType byte

const (
	Transient LifetimeType = iota
	Singleton
	Scoped
	Custom
)

func (t LifetimeType) String() string {
	switch t {
	case Transient:
		return "Transient"
	case Singleton:
		return "Singleton"
	case Scoped:
		return "Scoped"
	case Custom:
		return "Custom"
	}
	return "Unknown"
}

type LifetimeManagerBase interface {
	Instance(scopeId uuid.UUID) any
	SetInstance(instance any, scopeId uuid.UUID)
	ClearInstance(scopeId uuid.UUID)
}

type LifetimeManager interface {
	LifetimeManagerBase
	Type() LifetimeType
}
