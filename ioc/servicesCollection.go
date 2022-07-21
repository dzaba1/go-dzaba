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
		return nil, NewIocError(fmt.Sprintf("Invalid kind '%s'. Expected '%s'.", kind.String(), reflect.Func.String()))
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
		return nil, NewIocError(fmt.Sprintf("Invalid num of out args. Expected maximum 2, got %d.", outArgsCount))
	}
	if outArgsCount == 2 {
		errorType := utils.TypeOfGeneric[error]()
		secondType := ctorType.Out(1)
		if !utils.IsOrImplements(secondType, errorType) {
			return nil, NewIocError(fmt.Sprintf("Invalid second argument type. Expected '%s', got '%s'.", errorType.String(), secondType.String()))
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

func AddTransientSelf[T any](services ServicesCollection, ctorFunc any) error {
	genType := utils.TypeOfGeneric[T]()

	return services.AddTransientSelf(genType, ctorFunc)
}
