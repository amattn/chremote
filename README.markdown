
wdc is a simple, minimal webdriver client designed to be embedded in other apps.

There is deliberate emphasis on the **minimal** aspect.  While this project is an implementation against the [W3C webdriver spec][], it is currently a _partial_ implimentation.

The initial implementation plans to support chromium based browsers with the --enable-automation flag on.

[W3C webdriver spec]:https://w3c.github.io/webdriver/

## Usage

TODO

### Launching Chrome with the Enable Automation flag

On Mac OS X (make sure you quit the Chrome app first):

    open -a "Google Chrome" --args --incognito --remote-debugging-port=9222 http://www.example.com

 


## TODO

- [ ] finish writing this README
    