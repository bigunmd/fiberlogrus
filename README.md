# Logrus logger middleware
Logger middleware for Fiber that logs HTTP request/response details.


Use your configured `logrus` logger instance or global logrus instance to handle logging in a structured way.

## Table of Contents
- [Signatures](#signatures)
- [Examples](#examples)

## Signatures
```go
func New(config ...Config) fiber.Handler
```
## Examples
Import required packages
```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)
```
### Default config
Using with a default config, it will call global logrus instance to log the requests
```go
app := fiber.New()

app.Use(fiberlogrus.New())
```
```go
// ConfigDefault is the default config
var ConfigDefault Config = Config{
	Logger: nil,
	Tags:   []string{TagMethod},
}
```
### Use logger instance and configure tags
```go
logger := logrus.New()
// you can also provide logger with a desired formatter
// logger.SetFormatter(&logrus.JSONFormatter{})

app.Use(
	fiberlogrus.New(
		fiberlogrus.Config{
			Logger: logger,
			Tags: []string{
				// add method field
				fiberlogrus.TagMethod,
				// add status field
				fiberlogrus.TagStatus,
				// add value from locals
				AttachKeyTag(TagLocals, "requestid"),
				// add certain header
				AttachKeyTag(TagReqHeader, "custom-header"),
			},
		},
	),
)
```
### All supported common tags example
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	app.Use(
		fiberlogrus.New(
			fiberlogrus.Config{
				Logger: logger,
				Tags: fiberlogrus.CommonTags,
			}))
	
	app.Get("/", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })
	logger.Fatal(f.Listen(":8080"))
}
```

### Supported tags
#### Common
```go
// Common Tags
const (
	// request referer
	TagReferer = "referer"
	// request protocol
	TagProtocol = "protocol"
	// request port
	TagPort = "port"
	// request ip
	TagIP = "ip"
	// request ips
	TagIPs = "ips"
	// request host
	TagHost = "host"
	// request path
	TagPath = "path"
	// request url
	TagURL = "url"
	// request user-agent
	TagUA = "ua"
	// request body
	TagBody = "body"
	// request body bytes length
	TagBytesReceived = "bytesReceived"
	// response bytes length
	TagBytesSent = "bytesSent"
	// request route
	TagRoute = "route"
	// response body
	TagResBody = "resBody"
	// request headers
	TagReqHeaders = "reqHeaders"
	// request query parameters
	TagQueryStringParams = "queryParams"
	// response status
	TagStatus = "status"
	// request method
	TagMethod = "method"
	// fiber process id
	TagPid = "pid"
	// request latency
	TagLatency = "latency"
)
```
#### Key
```go
// Key Tags
const (
	// request specified header
	TagReqHeader = "reqHeader"
	// response specified header
	TagRespHeader = "respHeader"
	// request specified query
	TagQuery = "query"
	// request specified form value
	TagForm = "form"
	// request specified cookie value
	TagCookie = "cookie"
	// request specified locals value
	TagLocals = "locals"
)
```