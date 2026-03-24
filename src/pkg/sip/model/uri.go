package model

type URI struct {
	Scheme string
	User   string
	Host   string
	Port   int
	Params map[string]string
}
