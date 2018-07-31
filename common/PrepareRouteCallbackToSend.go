package common

import (
	"vuvuzela.io/alpenhorn/log"
)

func PrepareRouteCallbackToSend(route, callback string) ([]byte, error) {
	b, err := GetByteFromTransferable(&TransferableRouteCallback{"callback_forward", RouteCallback{route, callback}})
	if err != nil {
		log.Error(err)
	}

	return b, err
}
