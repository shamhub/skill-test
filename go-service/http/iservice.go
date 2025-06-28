package http

import (
	"context"
	"net/http"
)

type IHttpClient interface {
	Get(ctx context.Context, requestDetail *HttpRequestDetail) (*response, error)
	GetWithHeaders(ctx context.Context, requestDetail *HttpRequestDetail) (*response, error)
}

type IResponse interface {
	GetHeader(key string) string
	Headers() http.Header
	GetBody() []byte
	GetStatus() int
	Bind(response []byte, i interface{}) error
}

type HttpRequestDetail struct {
	Api                 string
	QueryParams         map[string]interface{}
	Body                []byte
	AuthDetail          AuthDetail
	RequestLevelHeaders map[string]string
}

type AuthDetail struct {
	Token string
	Name  string
}
