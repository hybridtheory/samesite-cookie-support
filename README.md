# samesite-cookie-support [![CircleCI](https://circleci.com/gh/affectv/samesite-cookie-support.svg?style=svg)](https://circleci.com/gh/affectv/samesite-cookie-support)

Go library to detect if SameSite=None cookies are supported.

This code is an implementation of https://www.chromium.org/updates/same-site/incompatible-clients

It uses [uap-go](https://github.com/ua-parser/uap-go) library to parse the user agents and applies
the logic to determine if the browser supports SameSite=None cookies or not.

## Usage

### Install

```bash
go get github.com/affectv/samesite-cookie-support
```

### Example

```golang
package main

import (
    "fmt"
    uaparser "github.com/affectv/samesite-cookie-support"
)

func main() {
    uagent := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16"
    fmt.Println(uaparser.IsSameSiteCookieSupported(uagent)) // true

    uagent = "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_14_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16"
    fmt.Println(uaparser.IsSameSiteCookieSupported(uagent)) // false
}
```
