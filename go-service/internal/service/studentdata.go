package service

import (
	"github.com/shamhub/pdfprovider/internal/dao"
	"github.com/shamhub/pdfprovider/pkg/config"
	"github.com/shamhub/pdfprovider/types"
)

type pdfGenerationService struct {
	dataFetcher dao.GetStudentData
	apiKey      string
}

func NewPDFGenrationService() *pdfGenerationService {
	unidocConfig, err := config.NewUniDocCred()
	if err != nil {
		panic(err)
	}
	apiKey := unidocConfig.Get(types.UNIDOC_LICENSE_API_KEY)

	return &pdfGenerationService{
		dataFetcher: dao.NewDataFetcher(),
		apiKey:      apiKey,
	}
}
