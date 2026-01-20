package client

import (
	"net/http"
	"time"

	"httpyum/internal/parser"
)

type Response struct {
	StatusCode  int
	Status      string
	Headers     http.Header
	Body        []byte
	ContentType string
	Duration    time.Duration
	RequestTime time.Time
	Size        int64
}

type ExecutionResult struct {
	Request  *parser.Request
	Response *Response
	Error    error
	Success  bool
}
