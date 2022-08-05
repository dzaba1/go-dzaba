package async

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func someOperation(result int, err error) (int, error) {
	return result, err
}

func Test_RunWithResult_WhenFired_ThenICanWaitForResults(t *testing.T) {
	expected := 1

	sut := NewAsyncTaskWithResultAndError(func(cancellationToken CancellationToken) (int, error) {
		return someOperation(expected, nil)
	})

	sut.Run()
	result, err := sut.Wait()

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_RunRunWithResult_WhenFired_ThenICanCancelIt(t *testing.T) {
	started := make(chan bool)
	sut := NewAsyncTaskWithResultAndError(func(cancellationToken CancellationToken) (int, error) {
		started <- true
		for {
			select {
			case <-cancellationToken:
				return OnTaskWithResultCancelled[int]()
			default:
			}
		}
	})

	sut.Run()
	<-started

	sut.Cancel()
	result, err := sut.Wait()

	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
	assert.IsType(t, &CancelledError{}, err)
}
