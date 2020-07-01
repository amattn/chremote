package wdclib

import (
	"net/url"

	"github.com/amattn/deeperror"
)

func ItWorks() string {
	return "You've loaded and called a function from this library."
}

type Client struct {
	serverURL *url.URL //URL of the websocket server
}

func NewClient(serverURL *url.URL) *Client {
	client := Client{
		serverURL: serverURL,
	}
	return &client
}

func (c *Client) SomethingBad() error {
	derr := deeperror.New(3555146532, " failure:", nil)
	derr.AddDebugField("key", "value")
	return derr
}
