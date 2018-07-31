package common

import (
	"vuvuzela.io/alpenhorn/log"
	"net/http"
)

func PrepareRequestToSend(r *http.Request) ([]byte, error) {
	request := GoRequestToCallbackerRequest(r)
	marshalledRequest, err := GetByteFromTransferable(&TransferableRequest{"callback_forward", request})
	if err != nil {
		log.Error(err)
	}

	return marshalledRequest, err
}
