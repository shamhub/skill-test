package handlers

import (
	"context"
	"errors"

	"github.com/shamhub/pdfprovider/pkg/server"
)

func (p *PDFHandler) GetStudentReport(ctx context.Context) (string, error) {
	if !isPathValid(p.filePath) {
		return "", errors.New("file not found")
	}

	studentId := server.GetStudentId(ctx)
	err := p.studentDataService.CreateReport(studentId)
	if err != nil {
		return p.filePath, err
	}

	return p.filePath, nil
}
