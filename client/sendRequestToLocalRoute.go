package client

import (
	"github.com/YasnaTeam/callbacker/common"
	"fmt"
)

func sendRequestToLocalRoute(r *common.Request, notification func (title, text string)) {
	hc := NewHttpClient(r)
	err := hc.SendRequest()

	if err != nil {
		log.Errorf("There is an error on sending request to local route, %s.", err.Error())
	} else {
		log.Debugf("`%s` has been forward to local endpoint.", hc.request.Callback)
		route, err := routes.Get(hc.request.Callback)

		if err == nil {
			notification("Callbacker", fmt.Sprintf("Request has been sent to '%s'", route.(string)))
		}
	}
}
