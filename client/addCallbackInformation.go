package client

func addCallbackInformation(result, route string, notification func (title, text string)) {
	routes.Set(result, route)
	configuration.AddRouteCallback(route, result)

	if routes.Has(result) {
		log.Debugf("`%s` successfully added to route table.", result)
	}
}
