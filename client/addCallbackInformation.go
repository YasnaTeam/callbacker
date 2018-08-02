package client

func addCallbackInformation(result, route string) {
	routes.Set(result, route)
	configuration.AddRouteCallback(route, result)

	if routes.Has(result) {
		log.Debugf("`%s` successfully added to route table.", result)
	}
}
