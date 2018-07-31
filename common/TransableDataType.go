package common

import (
	"fmt"
	"encoding/json"
)

type Transferable interface {
	GetType() string
	GetCommand() string
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

func GetTransferableFromByte(d []byte) (Transferable, error) {
	var data *TransferableData = &TransferableData{}
	err := json.Unmarshal(d, data)
	if err != nil {
		return nil, &TransferableError{err.Error(), "general unmarshal"}
	}

	return convertRawByteToTransferable(data)
}

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

func convertRawByteToTransferable(data *TransferableData) (Transferable, error) {
	// on unmarshalling of an interface value, go converts an interface type to a map
	mappedValues := data.Data.(map[string]interface{})

	switch data.Type {
	case "string":
		return fillStringTransferableFromData(mappedValues), nil
	case "request":
		return fillRequestTransferableFromData(mappedValues), nil
	default:
		return nil, &TransferableError{"There is no unmarshaller present for " + data.Type, "general unmarshal"}
	}
}

func fillStringTransferableFromData(d map[string]interface{}) Transferable {
	var tString TransferableString = TransferableString{}

	tString.Command = d["command"].(string)
	tString.Data = d["data"].(string)

	return &tString
}

func fillRequestTransferableFromData(mappedValues map[string]interface{}) Transferable {
	var tRequest TransferableRequest = TransferableRequest{}

	request := mappedValues["data"].(Request)
	tRequest.Data = &request
	tRequest.Command = mappedValues["command"].(string)

	return &tRequest
}