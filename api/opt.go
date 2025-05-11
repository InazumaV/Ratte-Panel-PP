package api

import (
	"Ratte-Panel-PP/api/client/server"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"time"
)

type setAble interface {
	SetTimeout(timeout time.Duration)
}

func WithTimeout[t setAble](a t, timeout time.Duration) t {
	a.SetTimeout(timeout)
	return a
}

type RequestHooker struct {
	handle func(runtime.ClientRequest, strfmt.Registry) error
	runtime.ClientRequestWriter
}

func NewParamsHooker(
	p runtime.ClientRequestWriter,
	handle func(runtime.ClientRequest, strfmt.Registry) error,
) *RequestHooker {
	return &RequestHooker{
		ClientRequestWriter: p,
		handle:              handle,
	}
}

func (e *RequestHooker) WriteToRequest(r runtime.ClientRequest, r2 strfmt.Registry) error {
	if e.handle != nil {
		err := e.handle(r, r2)
		if err != nil {
			return err
		}
	}
	return e.ClientRequestWriter.WriteToRequest(r, r2)
}

type ResponseHooker struct {
	handle func(runtime.ClientResponse, runtime.Consumer) (interface{}, error)
	skip   bool
	runtime.ClientResponseReader
}

func NewResponseHooker(
	r runtime.ClientResponseReader,
	handle func(runtime.ClientResponse, runtime.Consumer) (interface{}, error),
) *ResponseHooker {
	return &ResponseHooker{
		handle:               handle,
		ClientResponseReader: r,
	}
}
func (e *ResponseHooker) ReadResponse(r runtime.ClientResponse, c runtime.Consumer) (interface{}, error) {
	if e.handle != nil {
		return e.handle(r, c)
	}
	return e.ClientResponseReader.ReadResponse(r, c)
}

func WithEtag(etag string) server.ClientOption {
	return func(o *runtime.ClientOperation) {
		if o == nil {
			return
		}
		o.Params = NewParamsHooker(
			o.Params,
			func(r runtime.ClientRequest, reg strfmt.Registry) error {
				if etag != "" {
					err := r.SetHeaderParam("If-None-Match", etag)
					if err != nil {
						return err
					}
				}
				return nil
			})
		o.Reader = NewResponseHooker(
			o.Reader,
			func(r runtime.ClientResponse, c runtime.Consumer) (interface{}, error) {
				if r.Code() == 304 {
					return nil, nil
				}
				return r, nil
			})
	}
}
