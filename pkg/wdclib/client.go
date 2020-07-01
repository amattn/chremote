package wdclib

import (
	"log"

	"github.com/amattn/deeperror"
	"golang.org/x/net/websocket"
)

// webdriver client
type Client struct {
	serverURL string //URL of the websocket server
	ws        *websocket.Conn
}

func NewClient(serverURL string) *Client {
	client := Client{
		serverURL: serverURL,
	}
	return &client
}

func (c *Client) Connect() error {

	var origin = "http://localhost/"
	ws, err := websocket.Dial(c.serverURL, "", origin)
	if err != nil {
		derr := deeperror.New(2824034815, " failure:", err)
		derr.AddDebugField("c", c)
		return derr
	}

	c.ws = ws

	if c.ws == nil {
		derr := deeperror.New(2824034816, "unexpected nil ws failure", err)
		derr.AddDebugField("c", c)
		return derr
	}

	return nil
}

func (c *Client) SendJSON(unmarshalledJSONPayload interface{}) error {
	if c.ws == nil {
		derr := deeperror.New(2354828376, "unexpected nil ws, did you forget to Connect()?", nil)
		derr.AddDebugField("c", c)
		return derr
	}

	log.Println(4284546107, c)
	log.Println(4284546108, c.ws)
	log.Println(4284546109, unmarshalledJSONPayload)
	err := websocket.JSON.Send(c.ws, unmarshalledJSONPayload)
	if err != nil {
		derr := deeperror.New(4196813969, "websocket.JSON.Send failure:", err)
		derr.AddDebugField("c", c)
		derr.AddDebugField("unmarshalledJSONPayload", unmarshalledJSONPayload)
		return derr
	}
	return nil
}
