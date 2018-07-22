package storage

type RouteTable interface {
	// Get routing information for a request. After receive a request, it map the
	// request to the client which must serve it.
	Get(route string) (string, error)

	// Set request information of a request. Route is the endpoint which received to
	// the server and value is information of client which registered this route
	Set(route, value string) error

	// Unset a route
	Unset(route string) error

	// Indicate user can set more route or not
	CanSet() bool
}

type RouteError struct {
	route string
}

func (e *RouteError) Error() string {
	return "Route " + e.route + " not found!"
}

type CanNotSetError struct {
	reason string
}

func (e *CanNotSetError) Error() string {
	return e.reason
}