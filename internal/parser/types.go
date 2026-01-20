package parser

type Request struct {
	ID          string
	LineStart   int
	LineEnd     int
	Method      string
	URL         string
	Headers     map[string]string
	Body        string
	Description string
}

type Variable struct {
	Name    string
	Value   string
	LineNum int
}

type ParsedFile struct {
	Variables []Variable
	Requests  []Request
	RawLines  []string
}
