package utils

import "reflect"

func TypeOfGeneric[T any]() reflect.Type {
	return reflect.TypeOf(DefaultGeneric[T]())
}

func DefaultGeneric[T any]() T {
	var empty T
	return empty
}

func IsOrImplements(currentType reflect.Type, expected reflect.Type) bool {
	return currentType == expected || expected.Kind() == reflect.Interface && currentType.Implements(expected)
}
