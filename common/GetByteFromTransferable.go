package common

import "encoding/json"

func GetByteFromTransferable(d Transferable) ([]byte, error) {
	var data TransferableData = TransferableData{
		Type: d.GetType(),
		Data: d,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, &TransferableError{err.Error(), "marshal"}
	}

	return b, nil
}
