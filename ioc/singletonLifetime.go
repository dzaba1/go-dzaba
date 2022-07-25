package ioc

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

func (t *singletonLifetime) Instance() any {
	return t.instance
}

func (t *singletonLifetime) SetInstance(instance any) {
	t.instance = instance
}
