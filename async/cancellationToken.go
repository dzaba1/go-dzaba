package async

type CancellationToken chan bool

func NewCancellationToken() CancellationToken {
	return make(CancellationToken)
}

func (token CancellationToken) Cancel() {
	token <- true
}

func (token CancellationToken) Close() {
	close(token)
}
