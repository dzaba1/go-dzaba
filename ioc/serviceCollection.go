package ioc

import (
	"dzaba/go-dzaba/utils"
	"errors"
	"fmt"
	"reflect"
)

type ServiceCollection interface {
	BuildServiceProvder() (ServiceProvider, error)
	Registrations() map[reflect.Type]Registration

	AddTransientSelf(selfType reflect.Type, ctorFunc any) error
}

type serviceCollectionImpl struct {
	registrations map[reflect.Type]Registration
}

func NewServiceCollection() ServiceCollection {
	return &serviceCollectionImpl{
		registrations: make(map[reflect.Type]Registration),
	}
}

func (services *serviceCollectionImpl) Registrations() map[reflect.Type]Registration {
	return services.registrations
}

func (services *serviceCollectionImpl) BuildServiceProvder() (ServiceProvider, error) {
	return newServiceProvider()
}

func (services *serviceCollectionImpl) AddTransientSelf(selfType reflect.Type, ctorFunc any) error {
	ctorDescriptor, err := getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	if !utils.IsOrImplements(selfType, ctorDescriptor.outArgType) {
		errMsg := fmt.Sprintf("Invalid constructor out type. Expected '%s', got '%s'.", selfType.String(), ctorDescriptor.outArgType.String())
		return errors.New(errMsg)
	}

	lifetime := newTransientLifetimeManager()
	registration := newRegistration(ctorDescriptor, selfType, selfType, lifetime)
	services.registrations[selfType] = registration

	return nil
}

func getCtorDescriptor(ctorFunc any) (*ctorDescriptor, error) {
	ctorType := reflect.TypeOf(ctorFunc)
	kind := ctorType.Kind()

	if kind != reflect.Func {
		return nil, errors.New(fmt.Sprintf("Invalid kind '%s'. Expected '%s'.", kind.String(), reflect.Func.String()))
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
		return nil, errors.New(fmt.Sprintf("Invalid num of out args. Expected maximum 2, got %d.", outArgsCount))
	}
	if outArgsCount == 2 {
		errorType := utils.TypeOfGeneric[error]()
		secondType := ctorType.Out(1)
		if !utils.IsOrImplements(secondType, errorType) {
			return nil, errors.New(fmt.Sprintf("Invalid second argument type. Expected '%s', got '%s'.", errorType.String(), secondType.String()))
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
