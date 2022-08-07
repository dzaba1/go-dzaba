package ioc

import "github.com/google/uuid"

type customLifetime struct {
	impl LifetimeManagerBase
}

func newCustomLifetimeManager(impl LifetimeManagerBase) LifetimeManager {
	return &customLifetime{
		impl: impl,
	}
}

func (t *customLifetime) Type() LifetimeType {
	return Custom
}

func (t *customLifetime) Instance(scopeId uuid.UUID) any {
	return t.impl.Instance(scopeId)
}

func (t *customLifetime) SetInstance(instance any, scopeId uuid.UUID) {
	t.impl.SetInstance(instance, scopeId)
}

func (t *customLifetime) ClearInstance(scopeId uuid.UUID) {
	t.impl.ClearInstance((scopeId))
}
