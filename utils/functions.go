package utils

import (
	"reflect"
)

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
	if currentType == expected {
		return true
	}

	expectedKind := expected.Kind()
	if expectedKind == reflect.Interface && currentType.Implements(expected) {
		return true
	}

	currentKind := currentType.Kind()

	if currentKind == reflect.Struct {
		fieldsNum := currentType.NumField()
		for i := 0; i < fieldsNum; i++ {
			field := currentType.Field(i)
			fieldType := field.Type
			if field.Anonymous && IsOrImplements(fieldType, expected) {
				return true
			}
		}
	}

	return false
}
