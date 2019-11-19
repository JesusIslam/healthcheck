# Healthcheck
[![Build Status](https://travis-ci.org/JesusIslam/healthcheck.svg?branch=master)](https://travis-ci.org/JesusIslam/healthcheck)
[![GoDoc](https://godoc.org/github.com/JesusIslam/healthcheck?status.svg)](https://godoc.org/github.com/JesusIslam/healthcheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/JesusIslam/healthcheck)](https://goreportcard.com/report/github.com/JesusIslam/healthcheck)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FJesusIslam%2Fhealthcheck.svg?type=small)](https://app.fossa.io/projects/git%2Bgithub.com%2FJesusIslam%2Fhealthcheck?ref=badge_small)

Healthcheck helper for Gin to be able to easily got health checked thoroughly by Consul

# Usage
Download and install it:

`go get github.com/JesusIslam/healthcheck`

Import it in your code:

`import "github.com/JesusIslam/healthcheck`

Simplest default example:

```
package main

import (
    "time"

    "github.com/JesusIslam/healthcheck"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    timeout := 15 * time.Second
    h := healthcheck.New(nil, "myauthtoken", timeout)
    r.GET("/health", h.Handle)

    r.Run(":8080")
}
```
