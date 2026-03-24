package model

type ViaHeader struct {
	ProtocolName    string
	ProtocolVersion string
	Transport       string
	Host            string
	Port            int
	Params          map[string]string
}
