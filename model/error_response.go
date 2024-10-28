package model

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}

// ErrorResponseError Contains information about an error that occurred.
type ErrorResponseError struct {
	// Message describing an error.
	ErrorMsg string `json:"error_msg"`
	// Path at which an error occurred.
	ErrorPath string `json:"error_path,omitempty"`
	// Name of the error.
	ErrorName string `json:"error_name"`
}
