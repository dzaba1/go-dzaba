package async

type CancelledError struct {
}

func (myErr *CancelledError) Error() string {
	return "The operation was cancelled."
}
