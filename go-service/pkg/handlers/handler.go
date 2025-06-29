package handlers

import (
	"github.com/shamhub/pdfprovider/internal/service"
)

type PDFHandler struct {
	filePath           string
	studentDataService service.IGenerateReport
}

func NewPDfHandler(apiKey, filePath, backendURL, backendPort string) *PDFHandler {
	return &PDFHandler{
		studentDataService: service.NewPDFGenerationService(apiKey, filePath, backendURL, backendPort),
		filePath:           filePath,
	}
}
