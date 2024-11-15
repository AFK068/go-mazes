package domain

type InvalidInput struct {
	Message string
}

func (e *InvalidInput) Error() string {
	return e.Message
}
