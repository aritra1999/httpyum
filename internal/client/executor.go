package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"httpyum/internal/parser"
)

type Executor struct {
	client    *http.Client
	variables map[string]string
}

func NewExecutor(variables map[string]string) *Executor {
	return &Executor{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		variables: variables,
	}
}

func (e *Executor) Execute(req *parser.Request) *ExecutionResult {
	startTime := time.Now()

	url := parser.SubstituteVariables(req.URL, e.variables)

	var bodyReader io.Reader
	if req.Body != "" {
		body := parser.SubstituteVariables(req.Body, e.variables)
		bodyReader = strings.NewReader(body)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		return &ExecutionResult{
			Request: req,
			Error:   NewExecutionError(req.ID, "failed to create request", err),
			Success: false,
		}
	}

	for key, value := range req.Headers {
		substitutedValue := parser.SubstituteVariables(value, e.variables)
		httpReq.Header.Set(key, substitutedValue)
	}

	httpResp, err := e.client.Do(httpReq)
	if err != nil {
		duration := time.Since(startTime)
		return &ExecutionResult{
			Request: req,
			Error:   NewExecutionError(req.ID, "request failed", err),
			Success: false,
			Response: &Response{
				Duration:    duration,
				RequestTime: startTime,
			},
		}
	}
	defer httpResp.Body.Close()

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		duration := time.Since(startTime)
		return &ExecutionResult{
			Request: req,
			Error:   NewExecutionError(req.ID, "failed to read response body", err),
			Success: false,
			Response: &Response{
				StatusCode:  httpResp.StatusCode,
				Status:      httpResp.Status,
				Headers:     httpResp.Header,
				Duration:    duration,
				RequestTime: startTime,
			},
		}
	}

	duration := time.Since(startTime)

	contentType := httpResp.Header.Get("Content-Type")

	response := &Response{
		StatusCode:  httpResp.StatusCode,
		Status:      httpResp.Status,
		Headers:     httpResp.Header,
		Body:        bodyBytes,
		ContentType: contentType,
		Duration:    duration,
		RequestTime: startTime,
		Size:        int64(len(bodyBytes)),
	}

	return &ExecutionResult{
		Request:  req,
		Response: response,
		Success:  true,
	}
}

func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func IsJSON(contentType string) bool {
	return strings.Contains(strings.ToLower(contentType), "json")
}

func PrettyPrintJSON(data []byte) (string, error) {
	var buf bytes.Buffer
	if err := indentJSON(&buf, data, "", "  "); err != nil {
		return string(data), err
	}
	return buf.String(), nil
}

func indentJSON(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	level := 0
	inString := false
	escape := false

	for i := 0; i < len(src); i++ {
		c := src[i]

		if escape {
			dst.WriteByte(c)
			escape = false
			continue
		}

		if c == '\\' && inString {
			dst.WriteByte(c)
			escape = true
			continue
		}

		if c == '"' {
			inString = !inString
			dst.WriteByte(c)
			continue
		}

		if inString {
			dst.WriteByte(c)
			continue
		}

		switch c {
		case '{', '[':
			dst.WriteByte(c)
			level++
			if i+1 < len(src) && src[i+1] != '}' && src[i+1] != ']' {
				dst.WriteByte('\n')
				dst.WriteString(prefix)
				for j := 0; j < level; j++ {
					dst.WriteString(indent)
				}
			}
		case '}', ']':
			if i > 0 && src[i-1] != '{' && src[i-1] != '[' {
				level--
				dst.WriteByte('\n')
				dst.WriteString(prefix)
				for j := 0; j < level; j++ {
					dst.WriteString(indent)
				}
			} else {
				level--
			}
			dst.WriteByte(c)
		case ',':
			dst.WriteByte(c)
			dst.WriteByte('\n')
			dst.WriteString(prefix)
			for j := 0; j < level; j++ {
				dst.WriteString(indent)
			}
		case ':':
			dst.WriteByte(c)
			dst.WriteByte(' ')
		case ' ', '\t', '\n', '\r':
			continue
		default:
			dst.WriteByte(c)
		}
	}

	return nil
}
