package net

import "github.com/valyala/fasthttp"

// GetIP get the IP from the header.
// It will try to fetch the client IP forwarded from Cloudflare.
func GetIP(ctx *fasthttp.RequestCtx) string {
	forwardedIP := ctx.Request.Header.Peek("X-Real-IP")
	if len(forwardedIP) != 0 {
		return string(forwardedIP)
	}
	return ctx.RemoteIP().String()
}
