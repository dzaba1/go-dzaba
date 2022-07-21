package utils

import "reflect"

func TypeOfGeneric[T any]() reflect.Type {
	var empty T
	typ := reflect.TypeOf(empty)

	if typ == nil {
		hack := func() T {
			return DefaultGeneric[T]()
		}
		hackType := reflect.TypeOf(hack)
		typ = hackType.Out(0)
	}

	return typ
}

func DefaultGeneric[T any]() T {
	var empty T
	return empty
}

func IsOrImplements(currentType reflect.Type, expected reflect.Type) bool {
	expectedKind := expected.Kind()

	return currentType == expected || expectedKind == reflect.Interface && currentType.Implements(expected)
}
