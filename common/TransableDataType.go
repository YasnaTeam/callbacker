package common

import (
	"fmt"
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
