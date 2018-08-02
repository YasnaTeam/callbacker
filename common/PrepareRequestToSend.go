package common

import (
	"vuvuzela.io/alpenhorn/log"
	"net/http"
)

func PrepareRequestToSend(r *http.Request, domain string) ([]byte, error) {
	request := GetInformationOfRequest(r, domain)
	b, err := GetByteFromTransferable(&TransferableRequest{"callback_forward", request})
	if err != nil {
		log.Error(err)
	}

	return b, err
}
