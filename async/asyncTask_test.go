package async

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Run_WhenFired_ThenICanWaitForResults(t *testing.T) {
	sut := NewAsyncTaskWithError(func(cancellationToken CancellationToken) error {
		return nil
	})

	sut.Run()
	err := sut.Wait()

	assert.Nil(t, err)
}

func Test_Run_WhenFired_ThenICanCancelIt(t *testing.T) {
	started := make(chan bool)
	sut := NewAsyncTaskWithError(func(cancellationToken CancellationToken) error {
		started <- true
		for {
			select {
			case <-cancellationToken:
				return OnTaskCancelled()
			default:
			}
		}
	})

	sut.Run()
	<-started

	sut.Cancel()
	err := sut.Wait()

	assert.NotNil(t, err)
	assert.IsType(t, &CancelledError{}, err)
}
