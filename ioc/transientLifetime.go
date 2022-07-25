package ioc

type transientLifetime struct {
}

func newTransientLifetimeManager() LifetimeManager {
	return &transientLifetime{}
}

func (t *transientLifetime) Type() LifetimeType {
	return Transient
}

func (t *transientLifetime) Instance() any {
	return nil
}

func (t *transientLifetime) SetInstance(instance any) {
	// Don't cache any instance
}
