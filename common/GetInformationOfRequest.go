package common

import "net/http"

func GetInformationOfRequest(r *http.Request, domain string) *Request {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return &Request{
		Method:   r.Method,
		Callback: domain + "/" + r.URL.Path,
		URL:      r.URL.String(),
		Proto:    r.Proto,
		Header:   Header(r.Header),
		Body:     body,
		Host:     r.Host,
	}
}
