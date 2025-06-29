package handlers

import (
	"github.com/shamhub/pdfprovider/internal/service"
)

type PDFHandler struct {
	studentDataService service.GetPdfReport
}

func NewPDfHandler() *PDFHandler {

	return &PDFHandler{}
}
