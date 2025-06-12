package errs

type ErrorResponse struct {
	Message string   `json:"error"`
	Details []string `json:"details,omitempty"` // Optional field for additional error details
}

func (err ErrorResponse) Error() string {
	return err.Message
}
