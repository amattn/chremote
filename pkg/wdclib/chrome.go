package wdclib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/amattn/deeperror"
)

// example payload
//
//[
//	{
//		"description": "",
//		"devtoolsFrontendUrl": "/devtools/inspector.html?ws=localhost:9222/devtools/page/2D6AC07C3AE08C0E112ACA57A98E3EBF",
//		"id": "2D6AC07C3AE08C0E112ACA57A98E3EBF",
//		"title": "Matt Nunogawa / Dev, Entrepreneur, Writer",
//		"type": "page",
//		"url": "https://amattn.com/",
//		"webSocketDebuggerUrl": "ws://localhost:9222/devtools/page/2D6AC07C3AE08C0E112ACA57A98E3EBF"
//	}
//]

type ChromeWebPagePayload struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

func boostrapChrome(c *Client) error {
	//log.Println(2096286590, util.CurrentFunction())
	httpClient := &http.Client{}

	resp, err := httpClient.Get(c.bootstrapURL)
	if err != nil {
		derr := deeperror.New(2536509830, "browser boostrapURL (check --remote-debugging-port flag is set) failure:", err)
		derr.AddDebugField("c", c)
		return derr
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		derr := deeperror.New(1417314106, "ReadAll failure:", err)
		return derr
	}

	payload := []ChromeWebPagePayload{}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		derr := deeperror.New(2378951452, " failure:", err)
		derr.AddDebugField("body", string(body))
		return derr
	}

	//log.Printf("4030851949 %+v", payload)

	if len(payload) > 0 {
		c.websocketURL = payload[0].WebSocketDebuggerUrl
	} else {
		derr := deeperror.New(730977134, "unexpected empty payload:", err)
		derr.AddDebugField("c", c)
		return derr
	}

	return nil
}
