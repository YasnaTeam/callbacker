package common

import (
	"net/http"
	"strings"
)

func GetInformationOfRequest(r *http.Request, domain string) *Request {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return &Request{
		Method:   r.Method,
		Callback: generateCallback(domain, r.URL.Path),
		URL:      r.URL.String(),
		Proto:    r.Proto,
		Header:   Header(r.Header),
		Body:     body,
		Host:     r.Host,
	}
}

func generateCallback(domain, callback string) string {
	callback = strings.TrimLeft(callback, "/")
	callbackInformation := strings.Split(callback, "/")

	return domain + "/" + callbackInformation[1]
}
