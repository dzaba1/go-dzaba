package async

type AsyncWithErrorDelegate func(cancellationToken CancellationToken) error
type AsyncDelegate func(cancellationToken CancellationToken)

type asyncTaskBase struct {
	cancellationToken CancellationToken
}

type AsyncTask struct {
	asyncTaskBase
	action       AsyncWithErrorDelegate
	errorChannel chan error
}

func NewAsyncTaskWithError(action AsyncWithErrorDelegate) *AsyncTask {
	task := &AsyncTask{
		action:       action,
		errorChannel: make(chan error),
	}
	task.cancellationToken = NewCancellationToken()
	return task
}

func NewAsyncTask(action AsyncDelegate) *AsyncTask {
	var funcWrap AsyncWithErrorDelegate = func(cancellationToken CancellationToken) error {
		action(cancellationToken)
		return nil
	}
	return NewAsyncTaskWithError(funcWrap)
}

func (task *asyncTaskBase) Cancel() {
	task.cancellationToken.Cancel()
}

func (task *AsyncTask) Close() {
	task.cancellationToken.Close()
	close(task.errorChannel)
}

func (task *AsyncTask) Run() {
	go task.runInternal()
}

func (task *AsyncTask) runInternal() {
	err := task.action(task.cancellationToken)
	task.errorChannel <- err
}

func (task *AsyncTask) Wait() error {
	defer task.Close()
	for {
		select {
		case <-task.cancellationToken:
			return OnTaskCancelled()
		case err := <-task.errorChannel:
			return err
		}
	}
}

func OnTaskCancelled() error {
	return &CancelledError{}
}

func WhenAll(tasks []*AsyncTask) []error {
	errors := []error{}

	for _, task := range tasks {
		err := task.Wait()
		errors = append(errors, err)
	}

	return errors
}
