package http

import (
	"context"

	httpclient "github.com/shamhub/pdfprovider/internal/http"
	httpresponse "github.com/shamhub/pdfprovider/internal/http"
)

type response struct {
	*httpresponse.Response
}

type httpClient struct {
	*httpclient.HttpClientService
}

func NewHttpClient(hostportURL string) *httpClient {
	return &httpClient{
		HttpClientService: httpclient.NewHTTPServiceWithOptions(hostportURL),
	}
}

func (h *httpClient) Get(ctx context.Context, requestDetail *HttpRequestDetail) (*response, error) {
	resp, err := h.HttpClientService.Get(ctx, requestDetail.Api, requestDetail.QueryParams, requestDetail.AuthDetail.Token)

	return &response{
		Response: resp,
	}, err
}

func (h *httpClient) GetWithHeaders(ctx context.Context, requestDetail *HttpRequestDetail) (*response, error) {
	resp, err := h.HttpClientService.GetWithHeaders(ctx, requestDetail.Api, requestDetail.QueryParams, requestDetail.RequestLevelHeaders, requestDetail.AuthDetail.Token)

	return &response{
		Response: resp,
	}, err
}
