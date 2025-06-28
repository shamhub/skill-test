package handlers

import (
	"github.com/shamhub/pdfprovider/internal/service"
)

type PDFHandler struct {
	studentDataService service.GetStudentData
}

func NewPDfHandler() *PDFHandler {

	return &PDFHandler{}
}
