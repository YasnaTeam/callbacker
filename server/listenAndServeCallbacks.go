package server

import "net/http"

func listenAndServeCallbacks(address, domain string) {
	r := muxRouter(domain)

	log.Info("Server serves requests on " + address + "...")
	log.Fatal(http.ListenAndServe(address, r))
}
