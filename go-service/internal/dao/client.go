package dao

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/shamhub/pdfprovider/internal/http"
)

type DataFetcher struct {
	hostportUrl string
	client      *http.HttpClientService
}

func NewDataFetcher(backendURL, backendPort string) *DataFetcher {
	hostportUrl := "http://" + backendURL + ":" + backendPort
	return &DataFetcher{
		hostportUrl: hostportUrl,
		client:      http.NewHTTPServiceWithOptions(hostportUrl),
	}
}

func (d *DataFetcher) GetStudentData(studentId string) (*StudentData, error) {
	api := "/api/v1/students/" + studentId

	resp, err := d.client.Get(context.Background(), api, nil, "")
	if err != nil {
		return nil, err
	}

	studentData := StudentData{}
	err = resp.Bind(resp.Body, &studentData)
	if err != nil {
		return nil, err
	}
	_ = json.NewDecoder(bytes.NewBuffer(resp.Body)).Decode(&studentData)
	return &studentData, nil
}
