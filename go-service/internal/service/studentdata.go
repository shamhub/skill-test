package service

import (
	"github.com/shamhub/pdfprovider/internal/dao"
	"github.com/shamhub/pdfprovider/pkg/config"
	"github.com/shamhub/pdfprovider/types"
	"github.com/unidoc/unipdf/v3/creator"
)

type pdfGenerationService struct {
	dataFetcher dao.GetStudentData
	apiKey      string
	filePath    string
	creator     *creator.Creator
}

func NewPDFGenerationService() *pdfGenerationService {
	unidocConfig, err := config.NewUniDocCred()
	if err != nil {
		panic(err)
	}
	apiKey := unidocConfig.Get(types.UNIDOC_LICENSE_API_KEY)
	filePath := unidocConfig.Get(types.FILE_PATH)

	if apiKey == "" || filePath == "" {
		panic("invalid config in .env")
	}

	pdfCreator, err := NewPdfCreator()
	if err != nil {
		panic(err)
	}

	return &pdfGenerationService{
		dataFetcher: dao.NewDataFetcher(),
		apiKey:      apiKey,
		filePath:    filePath,
		creator:     pdfCreator,
	}
}
