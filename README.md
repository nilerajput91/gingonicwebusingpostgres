Gin Web Framework

Build Status codecov Go Report Card GoDoc Join the chat at https://gitter.im/gin-gonic/gin Sourcegraph Open Source Helpers Release TODOs

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

Installation

To install Gin package, you need to install Go and set your Go workspace first.

    The first need Go installed (version 1.11+ is required), then you can use the below Go command to install Gin.

$ go get -u github.com/gin-gonic/gin

    Import it in your code:

import "github.com/gin-gonic/gin"

    (Optional) Import net/http. This is required for example if using constants such as http.StatusOK.

import "net/http"
