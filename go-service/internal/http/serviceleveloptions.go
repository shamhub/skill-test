package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"
)

type ServiceLevelOptions struct {
	ServiceLevelHeaders map[string]string

	// TODO
	// *OAuthOption
	// *SurgeProtectorOption
}

func NewServiceLevelOption() *ServiceLevelOptions {
	return &ServiceLevelOptions{
		ServiceLevelHeaders: map[string]string{
			"Accept": "application/json,application/xml,text/plain",
		},
	}
}

func (s *ServiceLevelOptions) setServiceLevelHeaders(req *http.Request, body []byte) {
	if body == nil {
		return
	}

	contentType := "text/plain"
	var t interface{}

	err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&t)
	if err == nil {
		contentType = "application/json"
	}

	err = xml.NewDecoder(bytes.NewBuffer(body)).Decode(&t)
	if err == nil {
		contentType = "application/xml"
	}

	req.Header.Add("content-type", contentType)
}

func setTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     time.Duration(60) * time.Second,
		DisableKeepAlives:   false,
	}
}
