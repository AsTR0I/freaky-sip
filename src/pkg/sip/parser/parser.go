package parser

import (
	"bytes"
	siplog "freaky-sip/pkg/sip/log"
	"freaky-sip/pkg/sip/model"
)

type Parser struct {
	log siplog.Logger
}

func New(logger siplog.Logger) *Parser {
	if logger == nil {
		logger = siplog.NopLogger{}
	}

	return &Parser{
		log: logger,
	}
}

func (p *Parser) ParseMessage(raw []byte) (model.Message, error) {
	if len(raw) == 0 {
		return nil, &ParseError{
			Op:  "parse message",
			Err: ErrInvalidMessage,
		}
	}

	separator := []byte(CRLF + CRLF)

	i := bytes.Index(raw, separator)
	if i < 0 {
		return nil, &ParseError{
			Op:  "parse message",
			Err: ErrInvalidMessage,
		}
	}

	head := raw[:i]
	body := raw[i+len(separator):]
	_ = body
	// split message by CRLF
	lines := bytes.Split(head, []byte(CRLF))
	// check the start line
	if len(lines) == 0 || len(lines[0]) == 0 {
		return nil, &ParseError{
			Op:  "parse start-line",
			Err: ErrInvalidStartLine,
		}
	}

	startLine := string(lines[0])

	rawHeaders := make(model.Headers, 0, len(lines)-1)

	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}

		colon := bytes.IndexByte(line, ':')
		if colon < 0 {
			return nil, &ParseError{
				Op:   "parse header",
				Err:  ErrInvalidHeaderLine,
				Line: string(line),
			}
		}

		name := bytes.TrimSpace(line[:colon])
		value := bytes.TrimSpace(line[colon+1:])

		if len(name) == 0 {
			return nil, &ParseError{
				Op:   "parse headers",
				Err:  ErrInvalidHeaderLine,
				Line: string(line),
			}
		}

		rawHeaders.Add(string(name), string(value))
	}

	if isResponseStartLine(startLine) {
		return nil, nil
	}

	method, rawURI, version, err := parseRequestStartLine(startLine)
	if err != nil {
		return nil, err
	}

	if version != model.SIPVersion {
		return nil, &ParseError{
			Op:   "parse request start-line",
			Err:  ErrInvalidStartLine,
			Line: startLine,
		}
	}

	uri, err := ParseURI(rawURI)
	if err != nil {
		return nil, err
	}
	_ = uri

	_ = method

	return &model.Request{
		Method:     method,
		Version:    version,
		URI:        uri,
		RawHeaders: rawHeaders,
		Body:       body,
	}, nil

}
