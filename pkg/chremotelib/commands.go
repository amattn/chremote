package chremotelib

import "github.com/amattn/deeperror"

// chrome remote dev tools spec cna be found here:
// https://chromedevtools.github.io/devtools-protocol/tot/Page/
//
// The commands & terminology here are loosely modeled after the command speced in W3C WebDriver Spec:
// https://w3c.github.io/webdriver/#navigate-to

func (c *Client) NavigateTo(url string) error {
	payload := map[string]interface{}{
		"method": "Page.navigate",
		"id":     2,
		"params": map[string]interface{}{
			"url":            url,
			"transitionType": "typed",
		},
	}

	err := c.SendJSON(payload)
	if err != nil {
		derr := deeperror.New(2303596588, "NavigateTo failure:", err)
		derr.AddDebugField("payload", payload)
		return derr
	}

	return nil
}
