package common

func fillRequestTransferableFromData(b []byte) (*TransferableData, error) {
	var tRequest TransferableRequest

	td, err := UnmarshalTransferable(&tRequest, b)
	if err != nil {
		return nil, err
	}

	return td, nil
}
