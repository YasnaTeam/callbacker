package common

type RouteCallback struct {
	Route    string `json:"route"`
	Callback string `json:"callback"`
}

type TransferableRouteCallback struct {
	Command string        `json:"command"`
	Data    RouteCallback `json:"data"`
}

func (s *TransferableRouteCallback) GetType() string {
	return "route_callback"
}

func (s *TransferableRouteCallback) GetData() interface{} {
	return s.Data
}

func (s *TransferableRouteCallback) GetCommand() string {
	return s.Command
}
