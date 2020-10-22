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

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
// ignoreCache: if true, browser cache is ignored (as if the user pressed Shift+refresh).
func (c *Client) PageReload(ignoreCache bool) error {
	payload := map[string]interface{}{
		"method": "Page.reload",
		"id":     2,
		"params": map[string]interface{}{
			"ignoreCache": ignoreCache,
			// there is a second param here, scriptToEvaluateOnLoad, that is currently not implemented in this library
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

// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
// Close browser gracefully.
func (c *Client) Shutdown() error {
	payload := map[string]interface{}{
		"method": "Browser.close",
		"id":     2,
		"params": map[string]interface{}{},
	}

	err := c.SendJSON(payload)
	if err != nil {
		derr := deeperror.New(2245867873, "Shutdown failure:", err)
		derr.AddDebugField("payload", payload)
		return derr
	}

	return nil
}
