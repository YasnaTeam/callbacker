package common

import "net/http"

func GoRequestToCallbackerRequest(r *http.Request) *Request {
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	return &Request{
		Method: r.Method,
		URL: r.URL.String(),
		Proto: r.Proto,
		Header: Header(r.Header),
		Body: body,
		Host: r.Host,
	}
}
