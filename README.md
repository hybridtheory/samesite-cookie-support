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

## Parser

This parser is based on regular expressions and tries to aim specifically to those browsers
we are interested in because they don't support SameSite cookies fully.

The parser is focused in performance and relative accuracy.

### Performance

The library comes with a benchmark to check run times and place some tests around the expected
execution time.

| User-Agent        | Average (μs)      | Error (μs)|
| ----------------- |:-----------------:| ---------:|
| chrome            | 49.640            | ±49.590   |
| ucbrowser         | 63.360            | ±74.163   |
| iphone            | 37.720            | ±37.991   |
| safari            | 77.740            | ±24.546   |
| ubuntu            | 27.580            | ±41.001   |
| embedded browser  | 203.480           | ±64.915   |
| chrome version    | 51.950            | ±33.064   |
| iphone            | 37.720            | ±37.991   |

## Library Development and Testing

This library was developed using [ginkgo](https://github.com/onsi/ginkgo)
and [gomega](https://github.com/onsi/gomega) test frameworks.

To execute the tests:

```go
ginkgo ./...
```
