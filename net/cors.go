package net

import "github.com/valyala/fasthttp"

// PublicAPICORSWrapper handles all incoming requests and applies CORS headers to them, useful for public APIs.
// If the request type is OPTIONS the server will respond with fasthttp.StatusOK.
func PublicAPICORSWrapper(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "OPTIONS,POST,GET")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
		if ctx.Request.Header.IsOptions() {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}
		h(ctx)
	}
}
