package ioc

import (
	"dzaba/go-dzaba/collections"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type resolver interface {
	resolve(serviceType reflect.Type, scopeId uuid.UUID) (any, error)
	resolveRegistration(reg *registrationImpl, scopeId uuid.UUID) (any, error)
}

type resolverImpl struct {
	services map[reflect.Type][]*registrationImpl
}

func newResolver(services map[reflect.Type][]*registrationImpl) resolver {
	return &resolverImpl{
		services: services,
	}
}

func (r *resolverImpl) resolve(serviceType reflect.Type, scopeId uuid.UUID) (any, error) {
	chain := collections.NewStack[reflect.Type]()

	return r.resolveRecurse(serviceType, scopeId, chain)
}

func (r *resolverImpl) resolveRegistration(reg *registrationImpl, scopeId uuid.UUID) (any, error) {
	chain := collections.NewStack[reflect.Type]()

	return r.resolveRecurseRegistration(reg, scopeId, chain)
}

func (r *resolverImpl) resolveRecurseRegistration(reg *registrationImpl, scopeId uuid.UUID, chain *collections.Stack[reflect.Type]) (any, error) {
	chain.Push(reg.serviceType)

	instance := reg.lifetimeManager.Instance(scopeId)
	if instance != nil {
		return instance, nil
	}

	args := []any{}
	for _, argType := range reg.ctorDescriptor.inArgTypes {
		arg, err := r.resolveRecurse(argType, scopeId, chain)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	instance, err := reg.ctorDescriptor.activate(args...)
	if err != nil {
		return nil, err
	}

	reg.lifetimeManager.SetInstance(instance, scopeId)
	chain.Pop()
	return instance, nil
}

func (r *resolverImpl) resolveRecurse(serviceType reflect.Type, scopeId uuid.UUID, chain *collections.Stack[reflect.Type]) (any, error) {
	loop := collections.AnyMust(chain.GetList(), func(elem reflect.Type) bool {
		return serviceType == elem
	})

	if loop {
		return nil, fmt.Errorf("loop detected. Chain: %s", formatChain(chain))
	}

	regs, exist := r.services[serviceType]
	if !exist {
		serviceKind := serviceType.Kind()
		if serviceKind == reflect.Array || serviceKind == reflect.Slice {
			serviceElementType := serviceType.Elem()
			return r.resolveArray(serviceElementType, scopeId, chain)
		}

		return nil, fmt.Errorf("the service '%s' is not registered. Chain: %s", serviceType.String(), formatChain(chain))
	}

	reg := collections.Last(regs)
	return r.resolveRecurseRegistration(reg, scopeId, chain)
}

func (r *resolverImpl) resolveArray(serviceType reflect.Type, scopeId uuid.UUID, chain *collections.Stack[reflect.Type]) (any, error) {
	regs, exist := r.services[serviceType]
	if !exist {
		return nil, fmt.Errorf("the service '%s' is not registered. Chain: %s", serviceType.String(), formatChain(chain))
	}

	sliceType := reflect.SliceOf(serviceType)
	instancesValues := reflect.MakeSlice(sliceType, 0, len(regs))
	for _, reg := range regs {
		instance, err := r.resolveRecurseRegistration(reg, scopeId, chain)
		if err != nil {
			return nil, err
		}

		instanceValue := reflect.ValueOf(instance)
		instancesValues = reflect.Append(instancesValues, instanceValue)
	}

	return instancesValues.Interface(), nil
}

func formatChain(chain *collections.Stack[reflect.Type]) string {
	return ""
}
