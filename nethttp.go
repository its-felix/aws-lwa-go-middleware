package aws_lwa_go_middleware

import (
	"context"
	"github.com/its-felix/aws-lwa-go-middleware/types"
	"net/http"
)

func DecorateRequest(r *http.Request, hb func(http.Header, string)) (*http.Request, context.CancelFunc, error) {
	var err error
	var rawRequestCtx []byte
	if v := r.Header.Get(RequestContextHeaderName); v != "" {
		rawRequestCtx = []byte(v)
	}

	var lc *types.LambdaContext
	if v := r.Header.Get(LambdaContextHeaderName); v != "" {
		lc, err = ParseLambdaContext(v)
	}

	hb(r.Header, LambdaContextHeaderName)
	hb(r.Header, RequestContextHeaderName)

	ctx, cancel := WrapContext(r.Context(), rawRequestCtx, lc)
	r = r.WithContext(ctx)

	return r, cancel, err
}

func NetHTTPMiddleware(handler http.Handler, opts ...Option) http.Handler {
	var o options
	o.apply(opts...)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, cancel, err := DecorateRequest(r, o.headerBehavior)
		defer cancel()

		if err != nil {
			if err = o.errorBehavior(err); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		handler.ServeHTTP(w, req)
	})
}
