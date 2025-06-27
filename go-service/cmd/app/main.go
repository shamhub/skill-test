package main

import (
	"log"
	"net/http"
	"os" // Assuming the PDF is generated and saved to the file system
)

func main() {
	http.HandleFunc("/download-report", func(w http.ResponseWriter, r *http.Request) {
		// Assume the PDF file is located at "path/to/your/report.pdf"
		filePath := "./report.pdf"

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
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
