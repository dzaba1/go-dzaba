package ioc

import (
	"dzaba/go-dzaba/utils"
	"errors"
	"fmt"
	"reflect"
)

type ServiceCollection interface {
	BuildServiceProvder() (ServiceProvider, error)
	Registrations() map[reflect.Type][]Registration

	AddTransientSelf(selfType reflect.Type, ctorFunc any) error
	AddSingletonSelf(selfType reflect.Type, ctorFunc any) error
	AddScopedSelf(selfType reflect.Type, ctorFunc any) error
	AddCustomLifetimeSelf(selfType reflect.Type, lifetime LifetimeManagerBase, ctorFunc any) error
}

type serviceCollectionImpl struct {
	registrations map[reflect.Type][]*registrationImpl
}

func NewServiceCollection() ServiceCollection {
	return &serviceCollectionImpl{
		registrations: make(map[reflect.Type][]*registrationImpl),
	}
}

func (services *serviceCollectionImpl) Registrations() map[reflect.Type][]Registration {
	newDict := make(map[reflect.Type][]Registration)
	for key, value := range services.registrations {
		newArray := []Registration{}
		for _, reg := range value {
			newArray = append(newArray, reg)
		}

		newDict[key] = newArray
	}
	return newDict
}

func (services *serviceCollectionImpl) BuildServiceProvder() (ServiceProvider, error) {
	resolver := newResolver(services.registrations)
	return newServiceProvider(resolver, services.registrations)
}

func validateSelfType(selfType reflect.Type, ctorDescriptor *ctorDescriptor) error {
	if !utils.IsOrImplements(selfType, ctorDescriptor.outArgType) {
		errMsg := fmt.Sprintf("Invalid constructor out type. Expected '%s', got '%s'.", selfType.String(), ctorDescriptor.outArgType.String())
		return errors.New(errMsg)
	}

	return nil
}

func (services *serviceCollectionImpl) AddTransientSelf(selfType reflect.Type, ctorFunc any) error {
	ctorDescriptor, err := getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	err = validateSelfType(selfType, ctorDescriptor)
	if err != nil {
		return err
	}

	lifetime := newTransientLifetimeManager()
	registration := newRegistration(ctorDescriptor, selfType, selfType, lifetime)
	services.addRegistration(registration)

	return nil
}

func (services *serviceCollectionImpl) addRegistration(registration *registrationImpl) {
	regs := services.registrations[registration.serviceType]
	regs = append(regs, registration)
	services.registrations[registration.serviceType] = regs
}

func (services *serviceCollectionImpl) AddSingletonSelf(selfType reflect.Type, ctorFunc any) error {
	ctorDescriptor, err := getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	err = validateSelfType(selfType, ctorDescriptor)
	if err != nil {
		return err
	}

	lifetime := newSingletonLifetimeManager()
	registration := newRegistration(ctorDescriptor, selfType, selfType, lifetime)
	services.addRegistration(registration)

	return nil
}

func (services *serviceCollectionImpl) AddScopedSelf(selfType reflect.Type, ctorFunc any) error {
	ctorDescriptor, err := getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	err = validateSelfType(selfType, ctorDescriptor)
	if err != nil {
		return err
	}

	lifetime := newScopedLifetimeManager()
	registration := newRegistration(ctorDescriptor, selfType, selfType, lifetime)
	services.addRegistration(registration)

	return nil
}

func (services *serviceCollectionImpl) AddCustomLifetimeSelf(selfType reflect.Type, lifetime LifetimeManagerBase, ctorFunc any) error {
	ctorDescriptor, err := getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	err = validateSelfType(selfType, ctorDescriptor)
	if err != nil {
		return err
	}

	lifetimeInternal := newCustomLifetimeManager(lifetime)
	registration := newRegistration(ctorDescriptor, selfType, selfType, lifetimeInternal)
	services.addRegistration(registration)

	return nil
}

func getCtorDescriptor(ctorFunc any) (*ctorDescriptor, error) {
	ctorType := reflect.TypeOf(ctorFunc)
	kind := ctorType.Kind()

	if kind != reflect.Func {
		return nil, fmt.Errorf("invalid kind '%s'. Expected '%s'.", kind.String(), reflect.Func.String())
	}

	inArgsCount := ctorType.NumIn()
	inArgTypes := []reflect.Type{}

	for i := 0; i < inArgsCount; i++ {
		argType := ctorType.In(i)
		inArgTypes = append(inArgTypes, argType)
	}

	hasError := false
	outArgsCount := ctorType.NumOut()
	if outArgsCount > 2 {
		return nil, fmt.Errorf("invalid num of out args. Expected maximum 2, got %d.", outArgsCount)
	}
	if outArgsCount == 2 {
		errorType := utils.TypeOfGeneric[error]()
		secondType := ctorType.Out(1)
		if !utils.IsOrImplements(secondType, errorType) {
			return nil, fmt.Errorf("invalid second argument type. Expected '%s', got '%s'.", errorType.String(), secondType.String())
		}

		hasError = true
	}

	outType := ctorType.Out(0)

	return &ctorDescriptor{
		ctor:       ctorFunc,
		ctorType:   ctorType,
		inArgTypes: inArgTypes,
		outArgType: outType,
		hasError:   hasError,
	}, nil
}

func AddTransientSelf[T any](services ServiceCollection, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddTransientSelf(genType, ctorFunc)
}

func AddSingletonSelf[T any](services ServiceCollection, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddSingletonSelf(genType, ctorFunc)
}

func AddScopedSelf[T any](services ServiceCollection, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddScopedSelf(genType, ctorFunc)
}

func AddCustomLifetimeSelf[T any](services ServiceCollection, lifetime LifetimeManagerBase, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddCustomLifetimeSelf(genType, lifetime, ctorFunc)
}
