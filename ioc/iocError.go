package ioc

type iocError struct {
	msg string
}

func NewIocError(msg string) error {
	return &iocError{
		msg: msg,
	}
}

func (e *iocError) Error() string {
	return e.msg
}
