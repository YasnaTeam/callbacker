package common

func PrepareRouteCallbackInformation(route, callback string) ([]byte, error) {
	return PrepareRouteCallback(route, callback, "callback_information")
}
