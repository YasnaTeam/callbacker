package common

import (
	"vuvuzela.io/alpenhorn/log"
	"net/http"
)

func PrepareDataToSend(r *http.Request) ([]byte, error) {
	request := GoRequestToCallbackerRequest(r)
	marshalledRequest, err := GetTransferableDataByteFromInterface(&TransferableRequest{request}, "callback_forward")
	if err != nil {
		log.Error(err)
	}

	return marshalledRequest, err
}
