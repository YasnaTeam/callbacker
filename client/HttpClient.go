package client

import (
	"bytes"
	"net/http"

	"github.com/YasnaTeam/callbacker/common"
	"github.com/gojektech/heimdall/httpclient"
)

type Client struct {
	hc      *httpclient.Client
	request *common.Request
}

type Error struct {
	error string
}

func (c *Error) Error() string {
	return c.error
}

func NewHttpClient(r *common.Request) *Client {
	return &Client{httpclient.NewClient(), r}
}

func (c *Client) CreateRequest(method string) (*http.Request, error) {
	route, err := routes.Get(c.request.Callback)
	body := bytes.NewBuffer(c.request.Body)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, route.(string), body)

	c.setRequestHeader(req)

	return req, nil
}

func (c *Client) SendRequest() error {
	req, err := c.CreateRequest(c.request.Method)
	if err != nil {
		return err
	}

	_, err = c.hc.Do(req)

	return err
}

func (c *Client) setRequestHeader(r *http.Request) {
	for key, value := range c.request.Header {
		r.Header.Set(key, value[0])
	}
}
