package client

type errorString struct {
	Message string `json:"error_message"`
}

//Error override error print
func (e *errorString) Error() string {
	return e.Message
}

// NewError returns an error that formats as the given text.
func NewError(text string) error {
	return &errorString{text}
}
