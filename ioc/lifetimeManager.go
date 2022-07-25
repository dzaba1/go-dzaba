package ioc

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

type LifetimeManager interface {
	Type() LifetimeType
	Instance() any
	SetInstance(instance any)
}

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
