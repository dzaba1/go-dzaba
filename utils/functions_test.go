package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myInterface interface {
	Pass(msg string) string
}

type subMyInterface interface {
	myInterface
}

type myInterfaceImpl struct {
}

type subMyInterfaceImpl struct {
	myInterfaceImpl
}

type subSubMyInterfaceImpl struct {
	subMyInterfaceImpl
}

func NewMyInterface() myInterface {
	return &myInterfaceImpl{}
}

func (i *myInterfaceImpl) Pass(msg string) string {
	return msg
}

func Test_DefaultGeneric_WhenInterfaceProvided_ThenItIsNotNil(t *testing.T) {
	result := TypeOfGeneric[myInterface]()

	assert.NotNil(t, result)
}

func Test_IsOrImplements_WhenTypesProvided_ThenItMakesChecks(t *testing.T) {
	isOrImplementsCheck(t, TypeOfGeneric[*myInterfaceImpl](), TypeOfGeneric[myInterface](), true)
	isOrImplementsCheck(t, TypeOfGeneric[myInterfaceImpl](), TypeOfGeneric[myInterface](), false)
	isOrImplementsCheck(t, TypeOfGeneric[myInterface](), TypeOfGeneric[*myInterfaceImpl](), false)
	isOrImplementsCheck(t, TypeOfGeneric[int](), TypeOfGeneric[*int](), false)
	isOrImplementsCheck(t, TypeOfGeneric[*int](), TypeOfGeneric[int](), false)
	isOrImplementsCheck(t, TypeOfGeneric[int](), TypeOfGeneric[[]int](), false)
	isOrImplementsCheck(t, TypeOfGeneric[int](), TypeOfGeneric[map[int]string](), false)
	isOrImplementsCheck(t, TypeOfGeneric[subMyInterfaceImpl](), TypeOfGeneric[myInterfaceImpl](), true)
	isOrImplementsCheck(t, TypeOfGeneric[*subMyInterfaceImpl](), TypeOfGeneric[*myInterfaceImpl](), true)
	isOrImplementsCheck(t, TypeOfGeneric[*subMyInterfaceImpl](), TypeOfGeneric[myInterface](), true)
	isOrImplementsCheck(t, TypeOfGeneric[subMyInterface](), TypeOfGeneric[myInterface](), true)
	isOrImplementsCheck(t, TypeOfGeneric[subSubMyInterfaceImpl](), TypeOfGeneric[myInterfaceImpl](), true)
	isOrImplementsCheck(t, TypeOfGeneric[*subSubMyInterfaceImpl](), TypeOfGeneric[*myInterfaceImpl](), true)
	isOrImplementsCheck(t, TypeOfGeneric[*subSubMyInterfaceImpl](), TypeOfGeneric[myInterface](), true)
}

func isOrImplementsCheck(t *testing.T, current reflect.Type, toBe reflect.Type, expected bool) {
	result := IsOrImplements(current, current)
	assert.True(t, result)

	result = IsOrImplements(toBe, toBe)
	assert.True(t, result)

	result = IsOrImplements(current, toBe)
	assert.Equal(t, expected, result, "Current '%s', toBe '%s'", current.String(), toBe.String())
}
