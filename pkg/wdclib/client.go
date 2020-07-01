package wdclib

import (
	"io"
	"math/rand"

	"github.com/amattn/deeperror"
	"golang.org/x/net/websocket"
)

// webdriver client
type Client struct {
	serverURL string //URL of the websocket server
	ws        *websocket.Conn

	doneCh chan bool

	payloadHandler JSONPayloadHandler
	errorHandler   ReceiveErrorHandler
}

type JSONPayloadHandler func(tracer int64, payload interface{})
type ReceiveErrorHandler func(tracer int64, err error)

func NewClient(serverURL string, payloadHandler JSONPayloadHandler, errorHandler ReceiveErrorHandler) *Client {
	client := Client{
		serverURL:      serverURL,
		payloadHandler: payloadHandler,
		errorHandler:   errorHandler,
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

	c.doneCh = make(chan bool)

	return nil
}

func (c *Client) Listen() {
	for {
		select {

		// time to quit
		case <-c.doneCh:
			c.doneCh <- true
			return

		// read data from websocket connection
		default:
			tracer := rand.Int63()
			var payload interface{}
			err := websocket.JSON.Receive(c.ws, &payload)
			if c.payloadHandler != nil {
				go c.payloadHandler(tracer, payload)
			}
			if err == io.EOF {
				// we go disconnected...
				// time to quit
				c.doneCh <- true
			} else if err != nil {
				// we got some other unspecified error
				// send it along to someone who will look at it. (we hope).
				if c.errorHandler != nil {
					go c.errorHandler(tracer, err)
				}
			}
		}
	}
}

func (c *Client) SendJSON(unmarshalledJSONPayload interface{}) error {
	if c.ws == nil {
		derr := deeperror.New(2354828376, "unexpected nil ws, did you forget to Connect()?", nil)
		derr.AddDebugField("c", c)
		return derr
	}

	err := websocket.JSON.Send(c.ws, unmarshalledJSONPayload)
	if err != nil {
		derr := deeperror.New(4196813969, "websocket.JSON.Send failure:", err)
		derr.AddDebugField("c", c)
		derr.AddDebugField("unmarshalledJSONPayload", unmarshalledJSONPayload)
		return derr
	}
	return nil
}
