package parser

import (
	"reflect"
	"testing"
)

func TestParseURI(t *testing.T) {
	tests := []struct {
		name       string
		raw        string
		wantScheme string
		wantUser   string
		wantHost   string
		wantPort   int
		wantParams map[string]string
		wantErr    bool
	}{
		{
			name:       "sip user host",
			raw:        "sip:alice@example.com",
			wantScheme: "sip",
			wantUser:   "alice",
			wantHost:   "example.com",
			wantPort:   0,
			wantParams: map[string]string{},
		},
		{
			name:       "sips user host",
			raw:        "sips:alice@example.com",
			wantScheme: "sips",
			wantUser:   "alice",
			wantHost:   "example.com",
			wantPort:   0,
			wantParams: map[string]string{},
		},
		{
			name:       "sip user host port",
			raw:        "sip:alice@example.com:5060",
			wantScheme: "sip",
			wantUser:   "alice",
			wantHost:   "example.com",
			wantPort:   5060,
			wantParams: map[string]string{},
		},
		{
			name:       "sip user host params",
			raw:        "sip:alice@example.com;transport=udp;lr",
			wantScheme: "sip",
			wantUser:   "alice",
			wantHost:   "example.com",
			wantPort:   0,
			wantParams: map[string]string{
				"transport": "udp",
				"lr":        "",
			},
		},
		{
			name:       "sip user host port params",
			raw:        "sip:alice@example.com:5060;transport=udp;lr",
			wantScheme: "sip",
			wantUser:   "alice",
			wantHost:   "example.com",
			wantPort:   5060,
			wantParams: map[string]string{
				"transport": "udp",
				"lr":        "",
			},
		},
		{
			name:       "sip host only",
			raw:        "sip:example.com",
			wantScheme: "sip",
			wantUser:   "",
			wantHost:   "example.com",
			wantPort:   0,
			wantParams: map[string]string{},
		},
		// {
		// 	name:       "sip host only port",
		// 	raw:        "sip:example.com:5060",
		// 	wantScheme: "sip",
		// 	wantUser:   "",
		// 	wantHost:   "example.com",
		// 	wantPort:   5060,
		// 	wantParams: map[string]string{},
		// },
		// {
		// 	name:       "sip host only params",
		// 	raw:        "sip:example.com;transport=udp",
		// 	wantScheme: "sip",
		// 	wantUser:   "",
		// 	wantHost:   "example.com",
		// 	wantPort:   0,
		// 	wantParams: map[string]string{
		// 		"transport": "udp",
		// 	},
		// },
		// {
		// 	name:       "sip user ipv6",
		// 	raw:        "sip:alice@[2001:db8::1]",
		// 	wantScheme: "sip",
		// 	wantUser:   "alice",
		// 	wantHost:   "2001:db8::1",
		// 	wantPort:   0,
		// 	wantParams: map[string]string{},
		// },
		// {
		// 	name:       "sip user ipv6 port",
		// 	raw:        "sip:alice@[2001:db8::1]:5060",
		// 	wantScheme: "sip",
		// 	wantUser:   "alice",
		// 	wantHost:   "2001:db8::1",
		// 	wantPort:   5060,
		// 	wantParams: map[string]string{},
		// },
		// {
		// 	name:       "sip ipv6 host only",
		// 	raw:        "sip:[2001:db8::1]",
		// 	wantScheme: "sip",
		// 	wantUser:   "",
		// 	wantHost:   "2001:db8::1",
		// 	wantPort:   0,
		// 	wantParams: map[string]string{},
		// },
		// {
		// 	name:       "sip ipv6 host only port",
		// 	raw:        "sip:[2001:db8::1]:5060",
		// 	wantScheme: "sip",
		// 	wantUser:   "",
		// 	wantHost:   "2001:db8::1",
		// 	wantPort:   5060,
		// 	wantParams: map[string]string{},
		// },
		// {
		// 	name:    "empty",
		// 	raw:     "",
		// 	wantErr: true,
		// },
		// {
		// 	name:    "missing scheme",
		// 	raw:     "alice@example.com",
		// 	wantErr: true,
		// },
		// {
		// 	name:    "empty host",
		// 	raw:     "sip:alice@",
		// 	wantErr: true,
		// },
		// {
		// 	name:    "broken ipv6 bracket",
		// 	raw:     "sip:alice@[2001:db8::1:5060",
		// 	wantErr: true,
		// },
		// {
		// 	name:    "broken port",
		// 	raw:     "sip:alice@example.com:abc",
		// 	wantErr: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURI(tt.raw)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got == nil {
				t.Fatalf("expected uri, got nil")
			}

			if got.Scheme != tt.wantScheme {
				t.Fatalf("scheme: got %q want %q", got.Scheme, tt.wantScheme)
			}
			if got.User != tt.wantUser {
				t.Fatalf("user: got %q want %q", got.User, tt.wantUser)
			}
			if got.Host != tt.wantHost {
				t.Fatalf("host: got %q want %q", got.Host, tt.wantHost)
			}
			if got.Port != tt.wantPort {
				t.Fatalf("port: got %d want %d", got.Port, tt.wantPort)
			}
			if !reflect.DeepEqual(got.Params, tt.wantParams) {
				t.Fatalf("params: got %#v want %#v", got.Params, tt.wantParams)
			}
		})
	}
}
