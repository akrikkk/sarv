package handlers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleHTML(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("Error reading HTML file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Парсим multipart форму с ограничением размера 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Получаем файл из поля "myFile" (как указано в вашей форме)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "No file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Обрабатываем содержимое
	result := service.WriteResults(string(content))

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
