package common

import "vuvuzela.io/alpenhorn/log"

func PrepareReceivedData(d []byte) (Transferable, error) {
	data, err := GetTransferableFromByte(d)
	if err != nil {
		log.Debugf("There is an error on receiving data: `%s`", err)
		return nil, err
	}

	return data, nil
}
