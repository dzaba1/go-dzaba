package ioc

import "github.com/google/uuid"

type singletonLifetime struct {
	instance any
}

func newSingletonLifetimeManager() LifetimeManager {
	return &singletonLifetime{
		instance: nil,
	}
}

func (t *singletonLifetime) Type() LifetimeType {
	return Singleton
}

func (t *singletonLifetime) Instance(scopeId uuid.UUID) any {
	return t.instance
}

func (t *singletonLifetime) SetInstance(instance any, scopeId uuid.UUID) {
	t.instance = instance
}

func (t *singletonLifetime) ClearInstance(scopeId uuid.UUID) {
	t.instance = nil
}
