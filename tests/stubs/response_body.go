package stubs

import "io"

type ResponseBody struct {
	Reader io.Reader
}

func (r *ResponseBody) Read(p []byte) (int, error) {
	return r.Reader.Read(p)
}

func (r *ResponseBody) Close() error {
	return nil
}
