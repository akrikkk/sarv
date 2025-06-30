package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func convertMorse(text string) string {
	if isMorse(text) {
		return morse.ToText(text)
	}
	return morse.ToMorse(text)
}

func isMorse(text string) bool {
	for _, letter := range text {
		if letter != '.' && letter != '-' && letter != ' ' {
			return false
		}
	}
	return true
}

func WriteResults(text string) string {
	convertedText := convertMorse(text)
	fileName := "results/" + time.Now().UTC().String()

	err := os.Mkdir("results", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	if err := os.WriteFile(fileName, []byte(convertedText), 0755); err != nil {
		log.Fatal(err)
	}
	return convertedText
}
