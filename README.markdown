
chremote is a simple, minimal Chrome DevTools client designed to be embedded in other apps.

There is deliberate emphasis on the **minimal** aspect.  This project is intended to implement a subset of the expansive [Chrome DevTools Protocol][].

The initial implementation plans to support chromium based browsers with the --enable-automation flag on.

[Chrome DevTools Protocol]:https://chromedevtools.github.io/devtools-protocol/


## A short history of WebDriver and remote control of Chrome 

Selenium was originally a browser extension and IDE (I think?) that exposed a REST API that allowed remote control of a browser.  WebDriver was a different project meant to be a successor to Selenium. Later both tools merged into a single project called Selenium WebDriver.

At some point a modified version of that REST API became a W3C standard somewhat confusingly called WebDriver.   W3C’s WebDriver is basically a standard cross platform protocol for remote control of a browser.  [W3C webdriver spec][]

[W3C webdriver spec]:https://w3c.github.io/webdriver/

Originally, each browser had a different method of being remote controlled.  the different incarnations of Sellenium/WebDriver were a layer above that, essentially unifying the remote browser control interface.

With Chrome, you need expose the remote debugging port via launch flag or setting.  Chrome is remotely controlled via the Chrome DevTools Protocol (JSON over websocket based protocol) https://chromedevtools.github.io/devtools-protocol/

You need the intermediary ChromeDriver which translates the WebDriver REST API calls into the Chrome DevTools Protocol which actually controlls the browser. https://chromedriver.chromium.org

Some recent browsers support (or are in the process of supporting) the W3C WebDriver protocol natively. Without the need for ChromeDriver style intermediaries. 

If you want to remotely control Chrome, you have a range of options from heaviest to lightest: 
- Use Selenium Webdriver family to tools/IDEs/CLIs etc. or similar tooling from vendors such as saucelabs.
- WebDriver clients that speak webdriver protocol + any intermediary translation layers (chromedriver) 
- open source or other tools/libraries that speak Chrome DevTools Protocol directly, (you lose cross platform support at this level)
- write directly to the Chrome remote-debugging-port websocket only commands you need (as compared to 3. above, CDTP has hundreds of commands and queries).

This project is the last of the above list.  A super-lightweight, partial implementation of the Chrome DevTools Protocol.

## Usage

Primary usage is to import the lib into your go project.  The cli in the `cmd` is essentially an example/test app.



### Launching Chrome with the Enable Automation flag

On Mac OS X (make sure you quit the Chrome app first):

    open -a "Google Chrome" --args --incognito --remote-debugging-port=9222 http://www.example.com

If the remote debugging port is open, the following command should return something sensible:

    curl -sg http://localhost:9222/json 


## License

MIT.  See the LICENSE file for the legalese.

## TODO

- [ ] finish writing this README
- [ ] using some exponential backoff library, attempt to reconnect if disconnected.
    
