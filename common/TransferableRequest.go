package common

type TransferableRequest struct {
	Command string   `json:"command"`
	Data    *Request `json:"data"`
}

func (s *TransferableRequest) GetType() string {
	return "request"
}

func (s *TransferableRequest) GetData() interface{} {
	return s.Data
}

func (s *TransferableRequest) GetCommand() string {
	return s.Command
}
