package errors

import "net/http"

// RestErr is the base structure for error responses
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError is used to create a RestErr informing a BadRequest and a message
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bar_request",
	}
}
