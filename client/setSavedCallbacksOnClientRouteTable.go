package client

func setSavedCallbacksOnClientRouteTable() {
	for _, v := range configuration.Routes {
		routes.Set(v.Callback, v.Route)
	}
}
