package ioc

import "reflect"

type ctorDescriptor struct {
	ctor       any
	ctorType   reflect.Type
	inArgTypes []reflect.Type
	outArgType reflect.Type
	hasError   bool
}

func (c *ctorDescriptor) activate(args ...any) (any, error) {
	return nil, nil
}
