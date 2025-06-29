package service

func (p *pdfGenerationService) CreateReport(studentId string) error {

	studentData, err := p.dataFetcher.GetStudentData(studentId)
	if err != nil {
		return err
	}
	err = p.GeneratePdf(*studentData)
	if err != nil {
		return err
	}
	return nil
}
