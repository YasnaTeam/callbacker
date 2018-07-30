package common

type Header map[string][]string

// It's a minified version of http.Request and documentation of fields can be found on corresponded struct
type Request struct {
	Method string
	URL    string
	Proto  string
	Header Header
	Body   []byte
	Host   string
}
