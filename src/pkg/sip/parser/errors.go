package parser

import "errors"

var (
	ErrInvalidMessage       = errors.New("invalid sip message")
	ErrInvalidStartLine     = errors.New("invalid sip start line")
	ErrInvalidHeaderLine    = errors.New("invalid sip header line")
	ErrInvalidContentLength = errors.New("invalid sip content-length")
	ErrInvalidURI           = errors.New("invalid sip uri")
)

type ParseError struct {
	Op   string
	Err  error
	Line string
}

func (e *ParseError) Error() string {
	if e == nil {
		return ""
	}
	if e.Line == "" {
		return e.Op + ": " + e.Err.Error()
	}
	return e.Op + ": " + e.Err.Error() + ": " + e.Line
}

func (e *ParseError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}
