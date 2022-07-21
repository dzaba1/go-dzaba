package ioc

import (
	"dzaba/go-dzaba/utils"
	"fmt"
	"reflect"
)

type ServicesCollection interface {
	BuildServiceProvder() (ServiceProvider, error)
	AddTransientSelf(selfType reflect.Type, ctorFunc any) error
}

type servicesCollectionImpl struct {
}

func NewServiceCollection() ServicesCollection {
	return &servicesCollectionImpl{}
}

func (services *servicesCollectionImpl) BuildServiceProvder() (ServiceProvider, error) {
	return newServiceProvider()
}

func (services *servicesCollectionImpl) AddTransientSelf(selfType reflect.Type, ctorFunc any) error {
	ctorDescriptor, err := services.getCtorDescriptor(ctorFunc)
	if err != nil {
		return err
	}

	if ctorDescriptor.outArgType != selfType {
		if !selfType.Implements(ctorDescriptor.outArgType) {
			errMsg := fmt.Sprintf("Invalid constructor out type. Expected '%s', got '%s'.", selfType.String(), ctorDescriptor.outArgType.String())
			return NewIocError(errMsg)
		}
	}

	return nil
}

func (services *servicesCollectionImpl) getCtorDescriptor(ctorFunc any) (*ctorDescriptor, error) {
	ctorType := reflect.TypeOf(ctorFunc)
	kind := ctorType.Kind()

	if kind != reflect.Func {
		return nil, NewIocError("Invalid kind.")
	}

	inArgsCount := ctorType.NumIn()
	inArgTypes := []reflect.Type{}

	for i := 0; i < inArgsCount; i++ {
		argType := ctorType.In(i)
		inArgTypes = append(inArgTypes, argType)
	}

	outArgsCount := ctorType.NumOut()
	if outArgsCount > 2 {
		return nil, NewIocError("Invalid num of out args.")
	}

	outType := ctorType.Out(0)

	return &ctorDescriptor{
		ctor:       ctorFunc,
		ctorType:   ctorType,
		inArgTypes: inArgTypes,
		outArgType: outType,
		hasError:   outArgsCount == 2,
	}, nil
}

func AddTransientSelf[T any](services ServicesCollection, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddTransientSelf(genType, ctorFunc)
}
