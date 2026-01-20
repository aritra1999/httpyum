package client

import "fmt"

type ExecutionError struct {
	RequestID string
	Message   string
	Cause     error
}

func (e *ExecutionError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("execution error for %s: %s (caused by: %v)", e.RequestID, e.Message, e.Cause)
	}
	return fmt.Sprintf("execution error for %s: %s", e.RequestID, e.Message)
}

func (e *ExecutionError) Unwrap() error {
	return e.Cause
}

func NewExecutionError(requestID, message string, cause error) *ExecutionError {
	return &ExecutionError{
		RequestID: requestID,
		Message:   message,
		Cause:     cause,
	}
}
