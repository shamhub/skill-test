package service

import (
	"github.com/shamhub/pdfprovider/internal/dao"
	"github.com/unidoc/unipdf/v3/creator"
)

type pdfGenerationService struct {
	dataFetcher *dao.DataFetcher
	apiKey      string
	filePath    string
	creator     *creator.Creator
}

func NewPDFGenerationService(apiKey, filePath, backendURL, backendPort string) *pdfGenerationService {
	pdfCreator, err := NewPdfCreator(apiKey)
	if err != nil {
		panic(err)
	}

	return &pdfGenerationService{
		dataFetcher: dao.NewDataFetcher(backendURL, backendPort),
		apiKey:      apiKey,
		filePath:    filePath,
		creator:     pdfCreator,
	}
}
