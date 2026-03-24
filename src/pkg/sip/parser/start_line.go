package parser

import "strings"

func isResponseStartLine(line string) bool {
	return strings.HasPrefix(line, "SIP/2.0 ")
}

func parseRequestStartLine(line string) (method string, rawURI string, version string, err error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		return "", "", "", &ParseError{
			Op:   "parse request start-line",
			Err:  ErrInvalidStartLine,
			Line: line,
		}
	}

	method = strings.TrimSpace(parts[0])
	rawURI = strings.TrimSpace(parts[1])
	version = strings.TrimSpace(parts[2])

	if method == "" || rawURI == "" || version == "" {
		return "", "", "", &ParseError{
			Op:   "parse request start-line",
			Err:  ErrInvalidStartLine,
			Line: line,
		}
	}

	return method, rawURI, version, nil
}
