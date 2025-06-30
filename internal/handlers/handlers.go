package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

//1040 - 1103
//65 - 122

func HandleHTML(res http.ResponseWriter, req *http.Request) {
	HTML, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	res.Write(HTML)
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	req.Body.Close()

	scanner := bufio.NewScanner(&buf)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
	}

	text := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#$!$") {
			break
		}
		text += scanner.Text()
	}

	if strings.Contains(text, "------WebKitFormBoundary") {
		text_splited := strings.Split(text, "------WebKitFormBoundary")
		fmt.Println(text_splited[0])
		res.Write([]byte(service.WriteResults(text_splited[0])))
	} else {
		text = text[:len(text)-64]
		fmt.Println(text)
		res.Write([]byte(service.WriteResults(text)))
	}
}
