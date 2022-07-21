package ioc

import "reflect"

type ctorDescriptor struct {
	ctor       any
	ctorType   reflect.Type
	inArgTypes []reflect.Type
	outArgType reflect.Type
	hasError   bool
}
