package chremotelib

import (
	"github.com/amattn/chremote/internal/util"
	"github.com/amattn/deeperror"
)

// Utility

// wrapper function for
func sendJSON(debugNum int64, debugMessage string, client *Client, unmarshalledJSONPayload map[string]interface{}, handler ResponseHandler) (uint64, error) {

	// since we are over websockets, we get responses async.  the response matching this request will have the same id
	id := generateUniqueId()

	client.RegisterHandler(id, handler)
	id, err := client.SendJSON(id, unmarshalledJSONPayload)
	if err != nil {
		derr := deeperror.New(debugNum, debugMessage, err)
		derr.AddDebugField("unmarshalledJSONPayload", unmarshalledJSONPayload)
		derr.AddDebugField("id", id)
		return id, derr
	}

	return id, nil

}

// ######
// #     #   ##    ####  ######
// #     #  #  #  #    # #
// ######  #    # #      #####
// #       ###### #  ### #
// #       #    # #    # #
// #       #    #  ####  ######
//

// chrome remote dev tools spec cna be found here:
// https://chromedevtools.github.io/devtools-protocol/tot/Page/
//
// The commands & terminology here are loosely modeled after the command speced in W3C WebDriver Spec:
// https://w3c.github.io/webdriver/#navigate-to

// url to navigate to and optional response handler
// the client will always send the response to c.payloadHandler if it exists.
// if you want a local handler, pass in a handler function here.
func (c *Client) NavigateTo(url string, handler ResponseHandler) (uint64, error) {
	return c.NavigateFrameTo(url, "", handler)
}

func (c *Client) NavigateFrameTo(url string, frameId string, handler ResponseHandler) (uint64, error) {

	payload := map[string]interface{}{
		"method": "Page.navigate",
	}

	if frameId == "" {
		payload["params"] = map[string]interface{}{
			"url":            url,
			"transitionType": "typed",
		}
	} else {
		payload["params"] = map[string]interface{}{
			"url":            url,
			"frameId":        frameId,
			"transitionType": "typed",
		}
	}

	return sendJSON(1464031878, util.CurrentFunction(), c, payload, handler)
}

// https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload
// ignoreCache: if true, browser cache is ignored (as if the user pressed Shift+refresh).
func (c *Client) PageReload(ignoreCache bool, handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Page.reload",
		"params": map[string]interface{}{
			"ignoreCache": ignoreCache,
			// there is a second param here, scriptToEvaluateOnLoad, that is currently not implemented in this library
		},
	}

	return sendJSON(3621462765, util.CurrentFunction(), c, payload, handler)
}

// ######
// #     # #####   ####  #    #  ####  ###### #####
// #     # #    # #    # #    # #      #      #    #
// ######  #    # #    # #    #  ####  #####  #    #
// #     # #####  #    # # ## #      # #      #####
// #     # #   #  #    # ##  ## #    # #      #   #
// ######  #    #  ####  #    #  ####  ###### #    #
//

// https://chromedevtools.github.io/devtools-protocol/tot/Browser/#method-close
// Close browser gracefully.
func (c *Client) Shutdown(handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Browser.close",
		"params": map[string]interface{}{},
	}

	return sendJSON(89261776, util.CurrentFunction(), c, payload, handler)
}

// #######
// #       #    # #    # #        ##   ##### #  ####  #    #
// #       ##  ## #    # #       #  #    #   # #    # ##   #
// #####   # ## # #    # #      #    #   #   # #    # # #  #
// #       #    # #    # #      ######   #   # #    # #  # #
// #       #    # #    # #      #    #   #   # #    # #   ##
// ####### #    #  ####  ###### #    #   #   #  ####  #    #
//

// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-resetPageScaleFactor
func (c *Client) EmulationResetPageScaleFactor(handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Emulation.resetPageScaleFactor",
		"params": map[string]interface{}{},
	}

	return sendJSON(63309281, util.CurrentFunction(), c, payload, handler)
}

// https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setPageScaleFactor
func (c *Client) EmulationSetPageScaleFactor(factor float32, handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Emulation.setPageScaleFactor",
		"params": map[string]interface{}{
			"pageScaleFactor": factor,
		},
	}

	return sendJSON(63309280, util.CurrentFunction(), c, payload, handler)
}

// #######
//    #      ##   #####   ####  ###### #####
//    #     #  #  #    # #    # #        #
//    #    #    # #    # #      #####    #
//    #    ###### #####  #  ### #        #
//    #    #    # #   #  #    # #        #
//    #    #    # #    #  ####  ######   #
//

func (c *Client) TargetActivateTarget(targetId string, handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Target.activateTarget",
		"params": map[string]interface{}{
			"targetId": targetId,
		},
	}

	return sendJSON(63309280, util.CurrentFunction(), c, payload, handler)
}

func (c *Client) TargetGetTargets(handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Target.getTargets",
		"params": map[string]interface{}{},
	}

	return sendJSON(63309280, util.CurrentFunction(), c, payload, handler)
}

func (c *Client) TargetSetDiscoverTargets(discover bool, handler ResponseHandler) (uint64, error) {
	payload := map[string]interface{}{
		"method": "Target.setDiscoverTargets",
		"params": map[string]interface{}{
			"discover": discover,
		},
	}

	return sendJSON(63309280, util.CurrentFunction(), c, payload, handler)
}
