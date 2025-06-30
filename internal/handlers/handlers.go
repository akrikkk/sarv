package handlers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleHTML(res http.ResponseWriter, req *http.Request) {
	HTML, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	res.Write(HTML)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename

	if len(filename) < 4 || filename[len(filename)-4:] != ".txt" {
		http.Error(w, "Only .txt files are allowed", http.StatusBadRequest)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		result := service.WriteResults(string(content))

		// Возвращаем результат
		w.Header().Set("Content-Type", "text/plain")
		if _, err := w.Write([]byte(result)); err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	}

}
