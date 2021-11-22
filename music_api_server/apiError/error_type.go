package apiError

type APIError struct {
	Code int
	Message string
}

func (err APIError)Error() string{
	return err.Message
}
