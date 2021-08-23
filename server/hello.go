package server

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handler `hello` has been ran.")
	//requestlog := logger.WithRequest(log, r)
	io.WriteString(w, "Hello World!")
}
