package common

import (
	"fmt"
	"encoding/json"
)

type TransferableData struct {
	DataType string      `json:"data_type"`
	Data     interface{} `json:"data"`
}

type TransferableError struct {
	error string
}

func (e *TransferableError) Error() string {
	return fmt.Sprintf("There is an error on {Get,Set}ting TrasferableData: `%s`", e.error)
}

func GetTransferableDataFromByte(d []byte) (*TransferableData, error) {
	var data TransferableData = TransferableData{}
	err := json.Unmarshal(d, data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GetTransferableDataByteFromInterface(d interface{}, t string) ([]byte, error) {
	data := TransferableData{
		t,
		d,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}
