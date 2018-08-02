package server

import (
	"net"
	uuid2 "github.com/satori/go.uuid"
	"github.com/YasnaTeam/callbacker/common"
)

func registerCallback(conn net.Conn, rc *common.RouteCallback, domain string) {
	log.Debugf("Trying to add `%s` as a route...", rc.Route)

	uuid := uuid2.NewV4().String()
	callback := domain + "/" + uuid
	log.Debugf("Callback url is `%s`.", callback)

	username, err := connections.GetKey(conn)
	if err != nil {
		log.Errorf("`%s` is not available now.", username)
		return
	}

	routes.Set(uuid, username)

	b, err := common.PrepareRouteCallbackInformation(rc.Route, callback)
	if err != nil {
		log.Errorf("There are some errors on preparing RouteCallback data byte: `%s`.", err.Error())
		return
	}

	_, err = common.SendDataToConnection(conn, b)
	if err != nil {
		log.Errorf("Could not send callback information to client, %s", err.Error())
		return
	}
}
