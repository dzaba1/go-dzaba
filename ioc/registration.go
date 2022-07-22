package ioc

import "reflect"

type Registration interface {
	ServiceType() reflect.Type
	ImplementationType() reflect.Type
	Lifetime() LifetimeType
}

type registrationImpl struct {
	ctorDescriptor  *ctorDescriptor
	serviceType     reflect.Type
	implType        reflect.Type
	lifetimeManager LifetimeManager
}

func newRegistration(ctorDescriptor *ctorDescriptor,
	serviceType reflect.Type,
	implType reflect.Type,
	lifetimeManager LifetimeManager) Registration {

	return &registrationImpl{
		ctorDescriptor:  ctorDescriptor,
		serviceType:     serviceType,
		implType:        implType,
		lifetimeManager: lifetimeManager,
	}
}

func (r *registrationImpl) ServiceType() reflect.Type {
	return r.serviceType
}

func (r *registrationImpl) ImplementationType() reflect.Type {
	return r.implType
}

func (r *registrationImpl) Lifetime() LifetimeType {
	return r.lifetimeManager.Type()
}
