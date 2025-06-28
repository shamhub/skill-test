package http

import (
	"context"
	"net/http"
	"time"
)

type HttpClientService struct {
	*http.Client
	hostPortURL          string
	serviceLevelOptions  *ServiceLevelOptions
	contentTypeSupported responseType
}

type responseType int

const (
	JSON responseType = iota
	XML
	TEXT
	UNSUPPORTED_RESPONSE_TYPE
)

func NewHTTPServiceWithOptions(hostPortURL string) *HttpClientService {
	createHttpClient := func() *http.Client {
		return &http.Client{
			Transport: NewLoggingRoundTripper(),
			Timeout:   time.Duration(100) * time.Second,
		}
	}

	return &HttpClientService{
		hostPortURL:         hostPortURL,
		Client:              createHttpClient(),
		serviceLevelOptions: NewServiceLevelOption(),
	}
}

func (h *HttpClientService) Get(ctx context.Context, api string, queryParams map[string]interface{},
	authToken string) (*Response, error) {
	return h.GetWithHeaders(ctx, api, queryParams, map[string]string{}, authToken)
}

func (h *HttpClientService) GetWithHeaders(ctx context.Context, api string, queryParams map[string]interface{},
	headers map[string]string, authToken string) (*Response, error) {
	return h.callService(ctx, http.MethodGet, api, queryParams, nil, headers, authToken)
}
