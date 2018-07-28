package storage

type RouteTable interface {
	// Get routing information for a request. After receive a request, it map the
	// request to the client which must serve it.
	Get(key string) (interface{}, error)

	// Set request information of a request. Route is the endpoint which received to
	// the server and value is information of client which registered this key
	Set(key string, value interface{}) error

	Has(key string) bool

	// Unset a key
	Unset(key string) error

	// Indicate user can set more key or not
	CanSet() bool

	// Return first key which has the value
	GetKey(value interface{}) (string, error)
}

type RouteError struct {
	key string
}

func (e *RouteError) Error() string {
	return "Route " + e.key + " not found!"
}

type CanNotSetError struct {
	reason string
}

func (e *CanNotSetError) Error() string {
	return e.reason
}