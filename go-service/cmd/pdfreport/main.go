package main

import (
	"net/http"

	"github.com/shamhub/pdfprovider/pkg/config"
	"github.com/shamhub/pdfprovider/pkg/handlers"
	"github.com/shamhub/pdfprovider/pkg/server"
	"github.com/shamhub/pdfprovider/types"
)

func main() {

	handler := getReportHandler()
	server.GETReport("/api/v1/students/{id}/report", handler.GetStudentReport)
	http.ListenAndServe(":3000", server.MuxRouter)
}

func getReportHandler() *handlers.PDFHandler {
	envConfig, err := config.NewEnvConfig()
	if err != nil {
		panic(err)
	}

	//1. Read api key required to use pdf creator api
	apiKey := envConfig.Get(types.UNIDOC_LICENSE_API_KEY)
	if apiKey == "" {
		panic("invalid api key set in config")
	}

	//2. Read filepath to store pdf report
	filePath := envConfig.Get(types.FILE_PATH)
	if filePath == "" {
		panic("set file path in config")
	}

	//3. Read nodeJS backend url
	backendURL := envConfig.Get(types.BACKEND_URL)
	if backendURL == "" {
		panic("set node backend url in config")
	}

	//4. Read nodeJS backend port
	backendPort := envConfig.Get(types.BACKEND_PORT)
	if backendPort == "" {
		panic("set node backend port in config")
	}

	return handlers.NewPDfHandler(apiKey, filePath, backendURL, backendPort)
}
