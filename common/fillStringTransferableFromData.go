package common

func fillStringTransferableFromData(b []byte) (*TransferableData, error) {
	var tString TransferableString

	td, err := UnmarshalTransferable(&tString, b)
	if err != nil {
		return nil, err
	}

	return td, nil
}
