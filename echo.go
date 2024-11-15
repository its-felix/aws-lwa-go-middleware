package aws_lwa_go_middleware

import (
	"github.com/labstack/echo/v4"
)

func EchoMiddleware(opts ...Option) echo.MiddlewareFunc {
	var o options
	o.apply(opts...)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req, cancel, err := DecorateRequest(c.Request(), o.headerBehavior)
			defer cancel()

			if err != nil {
				if err = o.errorBehavior(err); err != nil {
					return err
				}
			}

			c.SetRequest(req)

			return next(c)
		}
	}
}
