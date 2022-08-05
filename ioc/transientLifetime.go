package ioc

import "github.com/google/uuid"

type transientLifetime struct {
}

func newTransientLifetimeManager() LifetimeManager {
	return &transientLifetime{}
}

func (t *transientLifetime) Type() LifetimeType {
	return Transient
}

func (t *transientLifetime) Instance(scopeId uuid.UUID) any {
	return nil
}

func (t *transientLifetime) SetInstance(instance any, scopeId uuid.UUID) {
	// Don't cache any instance
}

func (t *transientLifetime) ClearInstance(scopeId uuid.UUID) {
	// Don't cache any instance
}
