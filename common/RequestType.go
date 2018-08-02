package common

type Header map[string][]string

// It's a minified version of http.Request and documentation of fields can be found on corresponded struct
type Request struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	URL      string `json:"url"`
	Proto    string `json:"protocol"`
	Header   Header `json:"header"`
	Body     []byte `json:"body"`
	Host     string `json:"host"`
}
