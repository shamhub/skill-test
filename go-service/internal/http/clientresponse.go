package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

type Response struct {
	Body       []byte
	StatusCode int
	headers    http.Header
}

func (r *Response) GetHeader(key string) string {
	if r != nil && r.headers != nil {
		return r.headers.Get(key)
	}
	return ""
}

func (r *Response) Headers() http.Header {
	return r.headers.Clone()
}

func (r *Response) GetBody() []byte {
	if r != nil {
		return r.Body
	}
	return nil
}

func (r *Response) GetStatus() int {
	if r == nil {
		return -1
	}
	return r.StatusCode
}

// Bind() takes response and binds it to i based on content-type
func (r *Response) Bind(responseBody []byte, i interface{}) error {

	var err error

	switch getResponseContentType(r.headers) {
	case XML:
		err = xml.NewDecoder(bytes.NewBuffer(responseBody)).Decode(&i)
	case TEXT:
		v, ok := i.(*string)
		if ok {
			*v = string(responseBody)
		}
	case JSON:
		err = json.NewDecoder(bytes.NewBuffer(responseBody)).Decode(&i)
	default:
		err = errors.New("unsupported response type")
	}
	return err
}

func getResponseContentType(header http.Header) responseType {
	switch header.Get("content-type") {
	case "application/xml":
		return XML
	case "text/plain":
		return TEXT
	case "application/json":
		return JSON
	default:
		return UNSUPPORTED_RESPONSE_TYPE
	}
}
