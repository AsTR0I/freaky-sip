package model

const SIPVersion = "SIP/2.0"

type Header struct {
	Name  string
	Value string
}

type Headers []Header

func (h *Headers) Add(name, value string) {
	*h = append(*h, Header{
		Name:  name,
		Value: value,
	})
}

type Message interface {
	isMessage()
}

type Request struct {
	Method  string
	Version string
	URI     *URI

	Via           []ViaHeader
	From          *AddressHeader
	To            *AddressHeader
	CallID        *CallIDHeader
	CSeq          *CSeqHeader
	Contact       []AddressHeader
	MaxForwards   *MaxForwardsHeader
	ContentLength *ContentLengthHeader
	ContentType   *ContentTypeHeader
	RawHeaders    []Header
	Body          []byte
}

func (*Request) isMessage() {}

type Response struct {
	Version    string
	StatusCode int
	Reason     string

	Via           []ViaHeader
	From          *AddressHeader
	To            *AddressHeader
	CallID        *CallIDHeader
	CSeq          *CSeqHeader
	Contact       []AddressHeader
	ContentLength *ContentLengthHeader
	ContentType   *ContentLengthHeader
	RawHeaders    []Header
	Body          []byte
}

func (*Response) isMessage() {}
