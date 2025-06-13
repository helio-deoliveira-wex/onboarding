package api_response

type ErrorResponse struct {
	Message string   `json:"error"`
	Details []string `json:"details,omitempty"` // Optional field for additional error details
}
