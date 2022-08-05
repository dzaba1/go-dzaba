package async

import "dzaba/go-dzaba/utils"

type AsyncWithResultAndErrorDelegate[T any] func(cancellationToken CancellationToken) (T, error)
type AsyncWithResultDelegate[T any] func(cancellationToken CancellationToken) T

type AsyncTaskWithResult[T any] struct {
	asyncTaskBase
	action        AsyncWithResultAndErrorDelegate[T]
	resultChannel chan *resultWithError[T]
}

type resultWithError[T any] struct {
	result T
	err    error
}

func NewAsyncTaskWithResultAndError[T any](action AsyncWithResultAndErrorDelegate[T]) *AsyncTaskWithResult[T] {
	task := &AsyncTaskWithResult[T]{
		action:        action,
		resultChannel: make(chan *resultWithError[T]),
	}
	task.cancellationToken = NewCancellationToken()
	return task
}

func NewAsyncTaskWithResult[T any](action AsyncWithResultDelegate[T]) *AsyncTaskWithResult[T] {
	var funcWrap AsyncWithResultAndErrorDelegate[T] = func(cancellationToken CancellationToken) (T, error) {
		r := action(cancellationToken)
		return r, nil
	}
	return NewAsyncTaskWithResultAndError(funcWrap)
}

func (task *AsyncTaskWithResult[T]) Close() {
	task.cancellationToken.Close()
	close(task.resultChannel)
}

func (task *AsyncTaskWithResult[T]) Run() {
	go task.runInternal()
}

func (task *AsyncTaskWithResult[T]) runInternal() {
	result, err := task.action(task.cancellationToken)
	resultInternal := &resultWithError[T]{
		result: result,
		err:    err,
	}
	task.resultChannel <- resultInternal
}

func (task *AsyncTaskWithResult[T]) Wait() (T, error) {
	defer task.Close()
	for {
		select {
		case <-task.cancellationToken:
			return OnTaskWithResultCancelled[T]()
		case res := <-task.resultChannel:
			return res.result, res.err
		}
	}
}

func OnTaskWithResultCancelled[T any]() (T, error) {
	return utils.DefaultGeneric[T](), &CancelledError{}
}
