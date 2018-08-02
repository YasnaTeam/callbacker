package client

func truncateRouteTableAndFile() {
	log.Debug("Truncate routes table.")
	routes.Truncate()

	configuration.Routes = []RouteTable{}
	configuration.Save()
}
