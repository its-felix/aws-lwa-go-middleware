package aws_lwa_go_middleware

import (
	"context"
	"encoding/json"
	"github.com/its-felix/aws-lwa-go-middleware/types"
	"time"
)

const (
	RequestContextHeaderName = "X-Amzn-Request-Context"
	LambdaContextHeaderName  = "X-Amzn-Lambda-Context"
)

type requestContextKey struct{}
type lambdaContextKey struct{}

func ParseLambdaContext(raw string) (*types.LambdaContext, error) {
	var ctx *types.LambdaContext
	return ctx, json.Unmarshal([]byte(raw), &ctx)
}

func WrapContext(ctx context.Context, rcRaw []byte, lc *types.LambdaContext) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc

	if rcRaw != nil {
		ctx = context.WithValue(ctx, requestContextKey{}, rcRaw)
	}

	if lc != nil {
		ctx = context.WithValue(ctx, lambdaContextKey{}, lc)

		if lc.Deadline > 0 {
			ctx, cancel = context.WithDeadline(ctx, time.Unix(0, lc.Deadline))
		}
	}

	if cancel == nil {
		ctx, cancel = context.WithCancel(ctx)
	}

	return ctx, cancel
}

func RawRequestContext(ctx context.Context) ([]byte, bool) {
	v := ctx.Value(requestContextKey{})
	if v == nil {
		return nil, false
	}

	b, ok := v.([]byte)
	return b, ok
}

func RequestContext(ctx context.Context, v any) bool {
	b, ok := RawRequestContext(ctx)
	if !ok {
		return false
	}

	return json.Unmarshal(b, &v) == nil
}

func LambdaContext(ctx context.Context) (*types.LambdaContext, bool) {
	v := ctx.Value(lambdaContextKey{})
	if v == nil {
		return nil, false
	}

	lc, ok := v.(*types.LambdaContext)
	return lc, ok
}
