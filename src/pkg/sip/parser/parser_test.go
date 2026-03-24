package parser

import (
	"freaky-sip/pkg/sip/model"
	"testing"
)

func TestParseMessage_Request_StartLineAndHeaders(t *testing.T) {
	p := New(nil)

	raw := []byte(
		"INVITE sip:bob@example.com SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP client.example.com;branch=z9hG4bK-1\r\n" +
			"Content-Length: 0\r\n" +
			"\r\n",
	)

	msg, err := p.ParseMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req, ok := msg.(*model.Request)
	if !ok {
		t.Fatalf("expected *model.Request, got %T", msg)
	}

	if req.Method != "INVITE" {
		t.Fatalf("expected method INVITE, got %q", req.Method)
	}

	if req.Version != "SIP/2.0" {
		t.Fatalf("expected version SIP/2.0, got %q", req.Version)
	}

	if len(req.RawHeaders) != 2 {
		t.Fatalf("expected 2 raw headers, got %d", len(req.RawHeaders))
	}
}
