package aws_lwa_go_middleware

import (
	"errors"
	"net/http"
)

type options struct {
	errorBehavior  func(err error) error
	headerBehavior func(headers http.Header, name string)
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}

	if o.headerBehavior == nil {
		o.headerBehavior = func(headers http.Header, name string) {}
	}

	if o.errorBehavior == nil {
		o.errorBehavior = func(err error) error {
			return err
		}
	}
}

type Option func(o *options)

func WithIgnoreError() Option {
	return func(o *options) {
		o.errorBehavior = func(err error) error {
			return nil
		}
	}
}

func WithMaskError() Option {
	return func(o *options) {
		o.errorBehavior = func(err error) error {
			return errors.New(http.StatusText(http.StatusInternalServerError))
		}
	}
}

func WithRemoveHeaders() Option {
	return func(o *options) {
		o.headerBehavior = func(headers http.Header, name string) {
			headers.Del(name)
		}
	}
}
