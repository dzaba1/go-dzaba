package ioc

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func funcWithoutErr(intArg int, strArg string) string {
	return fmt.Sprintf("%s: %d", strArg, intArg)
}

func funcWithErr(intArg int, strArg string, err error) (string, error) {
	return funcWithoutErr(intArg, strArg), err
}

func Test_Activate_WhenCtorWithoutErr_ThenNoErrors(t *testing.T) {
	sut, err := getCtorDescriptor(funcWithoutErr)
	assert.Nil(t, err)

	result, err := sut.activate(10, "MyStr")
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func Test_Activate_WhenCtorWithErr_ThenErrors(t *testing.T) {
	sut, err := getCtorDescriptor(funcWithErr)
	assert.Nil(t, err)

	testErr := errors.New("Test error")
	result, err := sut.activate(10, "MyStr", testErr)
	assert.Equal(t, testErr, err)
	assert.NotEmpty(t, result)
}
