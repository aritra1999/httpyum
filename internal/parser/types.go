package parser

type Header struct {
	Key   string
	Value string
}

type Request struct {
	ID          string
	LineStart   int
	LineEnd     int
	Method      string
	URL         string
	Headers     []Header
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
