package dao

import "github.com/shamhub/pdfprovider/pkg/http"

type FetchData struct {
	client http.IHttpClient
}

func NewDataFetcher() *FetchData {
	return &FetchData{}
}
