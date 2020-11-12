package chremotelib

import (
	"fmt"
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

	registeredHandlers map[uint64]ResponseHandler
}

type JSONPayloadHandler func(tracer int64, payload interface{})
type ReceiveErrorHandler func(tracer int64, err error)
type ResponseHandler func(id uint64, payload map[string]interface{})

func NewClient(browserType SupportedUserAgents, browserURL string, payloadHandler JSONPayloadHandler, errorHandler ReceiveErrorHandler) *Client {
	client := Client{
		browserType:        browserType,
		payloadHandler:     payloadHandler,
		errorHandler:       errorHandler,
		registeredHandlers: make(map[uint64]ResponseHandler),
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

				//log.Printf("42785403010, %T", payload)
				// check registered handlers to see if we care about nay of them...
				// 1. first check to see if we have a map...
				typedPayload, isExpectedType := payload.(map[string]interface{})
				if isExpectedType {
					id, idExists := typedPayload["id"]
					//log.Printf("42785403011, %T", id)
					if idExists {
						uint64Id, isExpectedUInt64 := id.(uint64)
						float64Id, isExpectedFloat64 := id.(float64)

						if isExpectedFloat64 {
							uint64Id = uint64(float64Id)
						}

						if isExpectedUInt64 || uint64Id != 0 {
							handler := c.registeredHandlers[uint64Id]
							if handler != nil {
								go handler(uint64Id, typedPayload)
								delete(c.registeredHandlers, uint64Id)
							}
						}
					}

				} else {
					if c.errorHandler != nil {
						go c.errorHandler(tracer, fmt.Errorf("2599283748 Unknown payload format received over websocket"))
					}
				}
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

func (c *Client) RegisterHandler(id uint64, handler ResponseHandler) {
	//log.Println(2590942173, util.CurrentFunction(), id, handler)
	if handler == nil {
		return
	}

	// TODO 1532680991
	// not thread-safe...  handlers could blow away other handlers...
	c.registeredHandlers[id] = handler
}

// Send arbitrary JSON.
// Chrome will probably complain if you have excessive fields:
// map[error:map[code:-32600 message:Message has property other than 'id', 'method', 'sessionId', 'params']]
func (c *Client) SendJSON(id uint64, unmarshalledJSONPayload map[string]interface{}) (uint64, error) {

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

// don't use this directly.   call generateRequestId() to get a usable id
var internalAtomicIdCounter uint64

// we use an atomic counter.  Other options are random numbers, but the counter works for now.
func generateUniqueId() uint64 {
	updatedId := atomic.AddUint64(&internalAtomicIdCounter, 1)
	return updatedId
}
