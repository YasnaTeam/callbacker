package common

func PrepareRouteCallback(route, callback, command string) ([]byte, error) {
	b, err := GetByteFromTransferable(&TransferableRouteCallback{command, &RouteCallback{route, callback}})

	return b, err
}