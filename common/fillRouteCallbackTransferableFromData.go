package common

func fillRouteCallbackTransferableFromData(b []byte) (*TransferableData, error) {
	var tRouteCallback TransferableRouteCallback

	td, err := UnmarshalTransferable(&tRouteCallback, b)
	if err != nil {
		return nil, err
	}

	return td, nil
}
