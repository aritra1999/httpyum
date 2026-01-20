package parser

import "fmt"

type ParseError struct {
	Line    int
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at line %d: %s", e.Line, e.Message)
}

func NewParseError(line int, message string) *ParseError {
	return &ParseError{
		Line:    line,
		Message: message,
	}
}
