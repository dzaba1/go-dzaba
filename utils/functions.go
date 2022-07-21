package utils

import "reflect"

func TypeOfGeneric[T any]() reflect.Type {
	return reflect.TypeOf(DefaultGeneric[T]())
}

func DefaultGeneric[T any]() T {
	var empty T
	return empty
}
