package server

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type FileHandler func(context.Context) (path string, err error)

func (f FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background() // consume middleware code here
	filePath, err := f(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // error handling can be better
		return
	}

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Set the headers to serve the file as a download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.pdf\"")

	// Serve the file
	http.ServeFile(w, r, filePath)
}

var MuxRouter = mux.NewRouter().StrictSlash(false)

func GETReport(path string, handler FileHandler) {
	MuxRouter.NewRoute().Methods("GET").Path(path).Handler(handler)
}
