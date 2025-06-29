package service

import "github.com/shamhub/pdfprovider/internal/dao"

type PDFGenerationService struct {
	studentData dao.GetStudentData
}
