# aws-lwa-go-middleware
Middleware to be used with Go Lambdas using https://github.com/awslabs/aws-lambda-web-adapter.

## What it does
When applied, the middleware will populate the context with metadata supplied by the AWS LWA layer:
- The RequestContext will be read from `X-Amzn-Request-Context` and made available using `RequestContext(ctx context.Context, v any)` and `RawRequestContext(ctx context.Context)`
- The LambdaContext will be read from `X-Amzn-Lambda-Context` and made available using `LambdaContext(ctx context.Context)`
  - If a LambdaContext was found, the context of the incoming request will have a deadline set according to the `LambdaContext.Deadline`

## Usage

```golang
package main

import (
  lwamw "github.com/its-felix/aws-lwa-go-middleware"
  "github.com/labstack/echo/v4"
  "net/http"
)

func exampleNetHTTP() {
	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		// your request handler 
	})
	
	// options are optional
	handler = lwamw.NetHTTPMiddleware(
		handler, 
		lwamw.WithIgnoreError(), // ignore json parsing errors
		lwamw.WithMaskError(), // returns a 500 and masks the actual error
		lwamw.WithRemoveHeaders(), // removes both headers from the request before passing to the next MW/handler
	)
}

func exampleEcho() {
	e := echo.New()
	
	// this should be a global middleware and (probably) applied prior to any other middleware
	e.Use(lwamw.EchoMiddleware(
      lwamw.WithIgnoreError(), // ignore json parsing errors
      lwamw.WithMaskError(), // returns a 500 and masks the actual error
      lwamw.WithRemoveHeaders(), // removes both headers from the request before passing to the next MW/handler
	))
}
```