package server

import "net/http"

func listenAndServeCallbacks(address string) {
	r := muxRouter()

	log.Info("Server serves requests on " + address + "...")
	log.Fatal(http.ListenAndServe(address, r))
}
