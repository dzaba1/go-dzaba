package ioc

import (
	"dzaba/go-dzaba/collections"
	"fmt"
	"reflect"
)

type resolver interface {
	resolve(serviceType reflect.Type) (any, error)
}

type resolverImpl struct {
	services map[reflect.Type]*registrationImpl
}

func newResolver(services map[reflect.Type]*registrationImpl) resolver {
	return &resolverImpl{
		services: services,
	}
}

func (r *resolverImpl) resolve(serviceType reflect.Type) (any, error) {
	chain := collections.NewStack[reflect.Type]()

	return r.resolveRecurse(serviceType, chain)
}

func (r *resolverImpl) resolveRecurse(serviceType reflect.Type, chain *collections.Stack[reflect.Type]) (any, error) {
	loop := collections.AnyMust(chain.GetList(), func(elem reflect.Type) bool {
		return serviceType == elem
	})

	if loop {
		return nil, fmt.Errorf("loop detected. Chain: %s", formatChain(chain))
	}

	chain.Push(serviceType)

	reg, exist := r.services[serviceType]
	if !exist {
		return nil, fmt.Errorf("the service '%s' is not registered. Chain: %s", serviceType.String(), formatChain(chain))
	}

	instance := reg.lifetimeManager.Instance()
	if instance != nil {
		return instance, nil
	}

	args := []any{}
	for _, argType := range reg.ctorDescriptor.inArgTypes {
		arg, err := r.resolveRecurse(argType, chain)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	instance, err := reg.ctorDescriptor.activate(args...)
	if err != nil {
		return nil, err
	}

	reg.lifetimeManager.SetInstance(instance)
	chain.Pop()
	return instance, nil
}

func formatChain(chain *collections.Stack[reflect.Type]) string {
	return ""
}
