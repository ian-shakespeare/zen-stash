package handlers_test

import (
	"errors"
	"net/http"
)

const DEFAULT_STATUS_CODE int = 200

type DummyResponseWriter struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

func NewResponseWriter() *DummyResponseWriter {
	return &DummyResponseWriter{
		StatusCode: DEFAULT_STATUS_CODE,
		Body:       nil,
		Headers:    http.Header{},
	}
}

func (d *DummyResponseWriter) Header() http.Header {
	return d.Headers
}

func (d *DummyResponseWriter) Write(b []byte) (int, error) {
	if d.Body != nil {
		return 0, errors.New("body has already been written")
	}

	d.Body = b
	return len(b), nil
}

func (d *DummyResponseWriter) WriteHeader(statusCode int) {
	d.StatusCode = statusCode
}

func (d *DummyResponseWriter) Reset() {
	d.StatusCode = DEFAULT_STATUS_CODE
	d.Body = nil
	d.Headers = http.Header{}
}
