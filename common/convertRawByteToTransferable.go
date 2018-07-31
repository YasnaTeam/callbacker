package common

func convertRawByteToTransferable(data *TransferableData, b []byte) (*TransferableData, error) {
	// on unmarshalling of an interface value, go converts an interface type to a map
	//mappedValues := data.Data.(map[string]interface{})

	switch data.Type {
	case "string":
		return fillStringTransferableFromData(b)
	case "request":
		return fillRequestTransferableFromData(b)
	case "route_callback":
		return fillRouteCallbackTransferableFromData(b)
	default:
		return nil, &TransferableError{"There is no unmarshaller present for " + data.Type, "general unmarshal"}
	}
}
