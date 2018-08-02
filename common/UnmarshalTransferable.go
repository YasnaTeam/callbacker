package common

import "encoding/json"

func UnmarshalTransferable(t Transferable, b []byte) (*TransferableData, error) {
	td := &TransferableData{Data: t}

	if err := json.Unmarshal(b, td); err != nil {
		return nil, err
	}

	return td, nil
}
