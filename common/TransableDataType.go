package common

import (
	"fmt"
	"encoding/json"
)

type Transferable interface {
	GetType() string
	GetData() interface{}
}

type TransferableData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TransferableError struct {
	error     string
	errorType string
}

func (e *TransferableError) Error() string {
	return fmt.Sprintf("There is an error on %sing TrasferableData: `%s`", e.errorType, e.error)
}

func GetTransferableDataFromByte(d []byte) (*TransferableData, error) {
	var data TransferableData
	err := json.Unmarshal(d, &data)

	if err != nil {
		return nil, &TransferableError{err.Error(), "general unmarshal"}
	}

	switch data.Type {
	case "string":
		var tString TransferableString
		if err := json.Unmarshal([]byte(data.Raw), &tString); err != nil {
			return nil, &TransferableError{err.Error(), "string unmarshal"}
		}
		data.Data = tString.Data
		return &data, nil
	case "request":
		var tRequest TransferableRequest
		if err := json.Unmarshal([]byte(data.Raw), &tRequest); err != nil {
			return nil, &TransferableError{err.Error(), "request unmarshal"}
		}
		data.Data = tRequest.Data
		return &data, nil
	default:
		return nil, &TransferableError{"There is no unmarshaller present for " + data.Type, "general unmarshal"}
	}

	return &data, nil
}

func GetTransferableDataByteFromInterface(d Transferable, t string) ([]byte, error) {
	var data TransferableData = TransferableData{
		Type: d.GetType(),
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, &TransferableError{err.Error(), "marshal"}
	}

	return b, nil
}
