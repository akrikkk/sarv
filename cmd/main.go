package main

import (
	"errors"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	err := os.Mkdir("logs", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	flog, err := os.OpenFile(`logs/server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()

	logger := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)

	HTTPServer := server.StartServer(logger)

	HTTPServer.ListenAndServe()
}
