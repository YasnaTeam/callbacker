package common

func PrepareRouteCallbackToSend(route, callback string) ([]byte, error) {
	return PrepareRouteCallback(route, callback, "forward_callback")
}
