package handlers

import (
	"github.com/shamhub/pdfprovider/internal/service"
)

type PDFHandler struct {
	studentDataService service.IGenerateReport
}

func NewPDfHandler() *PDFHandler {

	return &PDFHandler{
		studentDataService: service.NewPDFGenerationService(),
	}
}
