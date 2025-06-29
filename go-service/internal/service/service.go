package service

import "github.com/shamhub/pdfprovider/internal/dao"

type IGenerateReport interface {
	GeneratePdf(studentData dao.StudentData) error
}
