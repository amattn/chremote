package chremotelib

import (
	"io"
	"log"
	"math/rand"
	"sync/atomic"

	"github.com/amattn/deeperror"
	"golang.org/x/net/websocket"
)

// the WebDriver spec calls the app (chrome.app, firefox.app, etc.) User Agents.
// a User Agent will have zero or more active browsers (window or tab)
type SupportedUserAgents string

const (
	Unknown SupportedUserAgents = "Unknown"
	Chrome                      = "Chrome"

	// below browsers not yet supported
	//Firefox            = "Firefox"
)

// webdriver client
type Client struct {
	browserType  SupportedUserAgents
	bootstrapURL string // browsers remote debugging port/URL
	websocketURL string // URL of the websocket server
	ws           *websocket.Conn

	doneCh chan bool

	payloadHandler JSONPayloadHandler
	errorHandler   ReceiveErrorHandler

	sessionID string
}

type JSONPayloadHandler func(tracer int64, payload interface{})
type ReceiveErrorHandler func(tracer int64, err error)

func NewClient(browserType SupportedUserAgents, browserURL string, payloadHandler JSONPayloadHandler, errorHandler ReceiveErrorHandler) *Client {
	client := Client{
		browserType:    browserType,
		payloadHandler: payloadHandler,
		errorHandler:   errorHandler,
	}

	switch browserType {
	case Chrome:
		client.bootstrapURL = browserURL
	}
	return &client
}

func (c *Client) Connect() error {
	//log.Println(3805270842, util.CurrentFunction())

	switch c.browserType {
	case Chrome:
		err := BoostrapChrome(c)
		if err != nil {
			derr := deeperror.New(1715423597, "BoostrapChrome failure:", err)
			derr.AddDebugField("c", c)
			return derr
		}
	default:
		derr := deeperror.New(3625125231, " unknown or unsupported brower:", nil)
		derr.AddDebugField("c.browserType", c.browserType)
		return derr
	}

	var origin = "http://localhost/"
	ws, err := websocket.Dial(c.websocketURL, "", origin)
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

// you can use the cleverly named websocat like so if you need a test server:
//     websocat -s 9222
func (c *Client) Listen() error {

	log.Println("listening:", c.websocketURL)

	for {
		select {

		// time to quit
		case <-c.doneCh:
			c.doneCh <- true
			return nil

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

func (c *Client) SendJSON(unmarshalledJSONPayload map[string]interface{}) (uint64, error) {

	// since we are over websockets, we get responses async.  the response matching this request will have the same id
	id := generateUniqueId()
	unmarshalledJSONPayload["id"] = id

	if c.ws == nil {
		derr := deeperror.New(2354828376, "unexpected nil ws, did you forget to Connect()?", nil)
		derr.AddDebugField("c", c)
		return id, derr
	}

	err := websocket.JSON.Send(c.ws, unmarshalledJSONPayload)
	if err != nil {
		derr := deeperror.New(4196813969, "websocket.JSON.Send failure:", err)
		derr.AddDebugField("c", c)
		derr.AddDebugField("unmarshalledJSONPayload", unmarshalledJSONPayload)
		return id, derr
	}
	return id, nil
}

var internalAtomicIdCounter uint64 // don't use this directly.   call generateRequestId() to get a usable id

// we use an atomic counter.  Other options are random numbers, but the counter works for now.
func generateUniqueId() uint64 {
	updatedId := atomic.AddUint64(&internalAtomicIdCounter, 1)
	return updatedId
}
