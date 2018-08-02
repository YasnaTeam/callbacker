package client

import "github.com/YasnaTeam/callbacker/common"

func sendRequestToLocalRoute(r *common.Request) {
	hc := NewHttpClient(r)
	err := hc.SendRequest()

	if err != nil {
		log.Errorf("There is an error on sending request to local route, %s.", err.Error())
	}
}
