package common

type TransferableString struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

func (s *TransferableString) GetType() string {
	return "string"
}

func (s *TransferableString) GetData() interface{} {
	return s.Data
}

func (s *TransferableString) GetCommand() string {
	return s.Command
}
