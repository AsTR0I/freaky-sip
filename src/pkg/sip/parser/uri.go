package parser

import (
	"freaky-sip/pkg/sip/model"
	"strconv"
	"strings"
)

const (
	SIP  string = "sip"
	SIPS string = "sips"
)

func ParseURI(uri string) (*model.URI, error) {
	uri = strings.TrimSpace(uri)
	if uri == "" {
		return nil, &ParseError{
			Op:  "parse uri",
			Err: ErrInvalidURI,
		}
	}

	if strings.HasPrefix(uri, SIPSProtocol) {
		return parse(SIPS, uri[len(SIPSProtocol):])
	}

	if strings.HasPrefix(uri, SIPProtocol) {
		return parse(SIP, uri[len(SIPProtocol):])
	}

	scheme := ""
	user := ""
	host := ""
	port := 0
	params := make(map[string]string, 0)

	return &model.URI{
		Scheme: scheme,
		User:   user,
		Host:   host,
		Port:   port,
		Params: params,
	}, nil
}

func parse(protocol string, uri string) (*model.URI, error) {

	// sip:example.com
	result := &model.URI{
		Scheme: protocol,
		Port:   0,
		Params: make(map[string]string),
	}

	uri = strings.TrimPrefix(uri, ":")

	sobakaIndex := strings.Index(uri, "@")
	if sobakaIndex == -1 {
		result.User = ""
	}

	result.User = strings.TrimSpace(uri[:sobakaIndex])

	hostPart := strings.TrimSpace(uri[sobakaIndex+1:])

	parts := strings.Split(hostPart, ";")

	hostPort := parts[0]

	// parse params
	for _, param := range parts[1:] {

		kv := strings.SplitN(param, "=", 2)

		key := kv[0]

		value := ""
		if len(kv) == 2 {
			value = kv[1]
		}

		result.Params[key] = value
	}

	// parse host + port
	hostPortParts := strings.SplitN(hostPort, ":", 2)

	result.Host = hostPortParts[0]

	if len(hostPortParts) == 2 {
		p, err := strconv.Atoi(hostPortParts[1])
		if err != nil {
			return nil, ErrInvalidURI
		}

		result.Port = p
	}

	return result, nil
}
