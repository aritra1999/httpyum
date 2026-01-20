package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	variableRegex   = regexp.MustCompile(`^@(\w+)\s*=\s*(.+)$`)
	httpMethodRegex = regexp.MustCompile(`^(GET|POST|PUT|DELETE|PATCH|HEAD|OPTIONS|TRACE|CONNECT)\s+(.+)$`)
	headerRegex     = regexp.MustCompile(`^([\w-]+)\s*:\s*(.+)$`)
	separatorRegex  = regexp.MustCompile(`^###`)
	commentRegex    = regexp.MustCompile(`^\s*(#|//)(.*)$`)
)

func Parse(r io.Reader) (*ParsedFile, error) {
	scanner := bufio.NewScanner(r)
	result := &ParsedFile{
		Variables: []Variable{},
		Requests:  []Request{},
		RawLines:  []string{},
	}

	lineNum := 0
	var currentRequest *Request
	var lastComment string
	inBody := false
	bodyLines := []string{}

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		result.RawLines = append(result.RawLines, line)
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			if currentRequest != nil && !inBody {
				inBody = true
			} else if inBody {
				bodyLines = append(bodyLines, line)
			}
			continue
		}

		if separatorRegex.MatchString(trimmedLine) {
			if currentRequest != nil {
				currentRequest.Body = strings.Join(bodyLines, "\n")
				currentRequest.LineEnd = lineNum - 1
				result.Requests = append(result.Requests, *currentRequest)
			}

			currentRequest = nil
			lastComment = ""
			inBody = false
			bodyLines = []string{}

			parts := strings.SplitN(trimmedLine, "###", 2)
			if len(parts) > 1 {
				comment := strings.TrimSpace(parts[1])
				if comment != "" {
					lastComment = comment
				}
			}
			continue
		}

		if commentMatches := commentRegex.FindStringSubmatch(trimmedLine); commentMatches != nil {
			if currentRequest == nil {
				comment := strings.TrimSpace(commentMatches[2])
				if comment != "" {
					lastComment = comment
				}
			}
			continue
		}

		if varMatches := variableRegex.FindStringSubmatch(trimmedLine); varMatches != nil {
			result.Variables = append(result.Variables, Variable{
				Name:    varMatches[1],
				Value:   strings.TrimSpace(varMatches[2]),
				LineNum: lineNum,
			})
			continue
		}

		if inBody {
			bodyLines = append(bodyLines, line)
			continue
		}

		if httpMatches := httpMethodRegex.FindStringSubmatch(trimmedLine); httpMatches != nil {
			if currentRequest != nil {
				currentRequest.Body = strings.Join(bodyLines, "\n")
				currentRequest.LineEnd = lineNum - 1
				result.Requests = append(result.Requests, *currentRequest)
			}

			currentRequest = &Request{
				ID:          fmt.Sprintf("req-%d", len(result.Requests)+1),
				LineStart:   lineNum,
				Method:      httpMatches[1],
				URL:         strings.TrimSpace(httpMatches[2]),
				Headers:     make(map[string]string),
				Description: lastComment,
			}
			lastComment = ""
			inBody = false
			bodyLines = []string{}
			continue
		}

		if currentRequest == nil && (strings.HasPrefix(trimmedLine, "http://") || strings.HasPrefix(trimmedLine, "https://")) {
			currentRequest = &Request{
				ID:          fmt.Sprintf("req-%d", len(result.Requests)+1),
				LineStart:   lineNum,
				Method:      "GET",
				URL:         trimmedLine,
				Headers:     make(map[string]string),
				Description: lastComment,
			}
			lastComment = ""
			inBody = false
			bodyLines = []string{}
			continue
		}

		if currentRequest != nil && !inBody {
			if headerMatches := headerRegex.FindStringSubmatch(trimmedLine); headerMatches != nil {
				headerName := headerMatches[1]
				headerValue := strings.TrimSpace(headerMatches[2])
				currentRequest.Headers[headerName] = headerValue
				continue
			}
		}

		if currentRequest != nil {
			if !inBody {
				inBody = true
			}
			bodyLines = append(bodyLines, line)
		}
	}

	if currentRequest != nil {
		currentRequest.Body = strings.Join(bodyLines, "\n")
		currentRequest.LineEnd = lineNum
		result.Requests = append(result.Requests, *currentRequest)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return result, nil
}

func SubstituteVariables(text string, variables map[string]string) string {
	pattern := regexp.MustCompile(`\{\{(\w+)\}\}`)

	result := pattern.ReplaceAllStringFunc(text, func(match string) string {
		varName := match[2 : len(match)-2]

		if value, ok := variables[varName]; ok {
			return value
		}

		return match
	})

	return result
}

func BuildVariableMap(variables []Variable) map[string]string {
	m := make(map[string]string)
	for _, v := range variables {
		m[v.Name] = v.Value
	}
	return m
}
