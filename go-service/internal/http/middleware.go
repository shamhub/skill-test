package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	next http.RoundTripper
	log  *log.Logger
}

func NewLoggingRoundTripper() *loggingRoundTripper {
	logger := log.New(os.Stdout, "HTTP_CLIENT_LOG:", log.Ldate|log.Ltime)

	return &loggingRoundTripper{
		next: setTransport(),
		log:  logger,
	}
}

func (lrt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var requestMethod string
	var reqURL string

	if req != nil {
		requestMethod = req.Method
		reqURL = req.URL.String()
	}

	requestData := struct {
		RequestMethod string `json:"requestmethod"`
		RequestURL    string `json:"requesturl"`
	}{
		RequestMethod: requestMethod,
		RequestURL:    reqURL,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	lrt.log.Printf("HTTPRequestDetails: %s\n", jsonData)

	start := time.Now()

	resp, err := lrt.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	duration := time.Since(start).Seconds()

	responseData := struct {
		ResponseStatus string  `json:"responsestatus"`
		Duration       float64 `json:"duration"`
	}{
		ResponseStatus: resp.Status,
		Duration:       duration,
	}
	jsonData, err = json.Marshal(responseData)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	lrt.log.Printf("HTTPResponseDetails: %s\n", jsonData)
	return resp, nil
}
