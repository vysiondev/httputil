package net

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

// GetRoot tries to get the base URL of the server that is handling the request through the Host header.
// If the host is localhost, the protocol will be http.
func GetRoot(ctx *fasthttp.RequestCtx) string {
	protocol := "https"
	if strings.Contains(string(ctx.Request.Host()), "localhost:") {
		protocol = "http"
	}
	return fmt.Sprintf("%s://%s", protocol, ctx.Request.Host())
}


