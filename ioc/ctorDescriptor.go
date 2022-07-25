package ioc

import (
	"dzaba/go-dzaba/collections"
	"reflect"
)

type ctorDescriptor struct {
	ctor       any
	ctorType   reflect.Type
	inArgTypes []reflect.Type
	outArgType reflect.Type
	hasError   bool
}

func (c *ctorDescriptor) activate(args ...any) (any, error) {
	ctorValue := reflect.ValueOf(c.ctor)
	argValues := collections.SelectMust(args, func(arg any) reflect.Value {
		return reflect.ValueOf(arg)
	})

	resultValues := ctorValue.Call(argValues)
	result := collections.SelectMust(resultValues, func(elem reflect.Value) any {
		return elem.Interface()
	})

	if len(result) == 1 {
		return result[0], nil
	}

	return result[0], result[1].(error)
}
