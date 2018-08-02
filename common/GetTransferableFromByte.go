package common

import (
	"encoding/json"
)

func GetTransferableFromByte(d []byte) (*TransferableData, error) {
	var data *TransferableData = &TransferableData{}
	err := json.Unmarshal(d, data)
	if err != nil {
		return nil, &TransferableError{err.Error(), "general unmarshal"}
	}

	return convertRawByteToTransferable(data, d)
}
