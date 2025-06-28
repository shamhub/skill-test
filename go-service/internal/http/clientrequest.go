package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	errorType "github.com/shamhub/pdfprovider/error"
)

func (h *HttpClientService) callService(ctx context.Context, method, api string,
	queryParams map[string]interface{}, body []byte, requestLevelHeaders map[string]string,
	authToken string) (*Response, error) {

	req, err := h.createReq(ctx, method, api, queryParams, body, requestLevelHeaders, authToken)
	if err != nil {
		return nil, err
	}

	statusCode := 0
	resp, err := h.Do(req.WithContext(ctx))
	if resp != nil {
		statusCode = resp.StatusCode
	}

	if err != nil {
		return nil, err
	}

	h.contentTypeSupported = getResponseContentType(resp.Header)

	if h.contentTypeSupported == UNSUPPORTED_RESPONSE_TYPE {
		return nil, errorType.UnSupportedResponse{
			ResponseType: "UnSupportedResponseType",
		}
	}

	responseBody, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()

	response := Response{
		Body:       responseBody,
		StatusCode: statusCode,
		headers:    resp.Header,
	}
	return &response, err
}

func (h *HttpClientService) createReq(ctx context.Context, method, api string, queryParams map[string]interface{},
	body []byte, requestLevelHeaders map[string]string, authToken string) (*http.Request, error) {

	httpUrl := h.hostPortURL + "/" + api
	if api == "" {
		httpUrl = h.hostPortURL
	}

	// 1. handle req url
	reqURL, err := url.Parse(httpUrl)
	if err != nil {
		log.Fatalf("invalid url: %s with error: %s\n", httpUrl, err.Error())
		return nil, err
	}

	//2. create new http request
	req, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, errorType.FailedRequest{URL: h.hostPortURL, Err: err}
	}

	// 3. Add query params
	if (method == "GET" ||
		method == "POST" ||
		method == "PUT" ||
		method == "PATCH") && queryParams != nil {
		encodeQueryParameters(reqURL, queryParams)
		req.URL = reqURL
	}

	// 4. Add service level headers for a http request
	h.serviceLevelOptions.setServiceLevelHeaders(req, body)

	// 5. Add request level headers
	for k, v := range requestLevelHeaders {
		req.Header.Set(k, v)
	}

	// 6. Add authorization header
	if authToken != "" {
		req.Header.Add("Authorization", authToken)
	}

	return req, nil
}

func encodeQueryParameters(reqURL *url.URL, queryParams map[string]interface{}) {

	query := reqURL.Query()

	for k, v := range queryParams {
		switch vt := v.(type) {
		case []string:
			for _, val := range vt {
				query.Set(k, val)
			}
		default:
			query.Set(k, fmt.Sprintf("%v", v))
		}
	}
	reqURL.RawQuery = query.Encode()
}
