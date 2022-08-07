package ioc

import "github.com/google/uuid"

type scopedLifetime struct {
	instances map[uuid.UUID]any
}

func newScopedLifetimeManager() LifetimeManager {
	return &scopedLifetime{
		instances: make(map[uuid.UUID]any),
	}
}

func (t *scopedLifetime) Type() LifetimeType {
	return Scoped
}

func (t *scopedLifetime) Instance(scopeId uuid.UUID) any {
	instance, ok := t.instances[scopeId]
	if ok {
		return instance
	}

	return nil
}

func (t *scopedLifetime) SetInstance(instance any, scopeId uuid.UUID) {
	t.instances[scopeId] = instance
}

func (t *scopedLifetime) ClearInstance(scopeId uuid.UUID) {
	t.instances[scopeId] = nil
}
