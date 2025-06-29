package service

type IGenerateReport interface {
	CreateReport(studentId string) error
}
