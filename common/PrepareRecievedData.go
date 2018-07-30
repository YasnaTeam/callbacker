package common

import "vuvuzela.io/alpenhorn/log"

func PrepareReceivedData(d []byte) (*TransferableData, error) {
	data, err := GetTransferableDataFromByte(d)
	if err != nil {
		log.Debugf("There is an error on receiving data: `%s`", err)
		return nil, err
	}

	return data, nil
}
