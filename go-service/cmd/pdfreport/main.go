package main

import (
	"net/http"

	"github.com/shamhub/pdfprovider/handlers"
	"github.com/shamhub/pdfprovider/server"
)

func main() {
	handler := getReportHandler()
	server.GETReport("/api/v1/students/{id}/report", handler.GetStudentData)
	http.ListenAndServe(":3000", server.MuxRouter)
}

func getReportHandler() *handlers.PDFHandler {
	return handlers.NewPDfHandler()
}
