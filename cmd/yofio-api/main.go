package main

import (
	"net/http"
	"os"

	"github.com/avelinoschz/yofio/internal/api"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

func main() {
	if err := run(); err != nil {
		log.Errorln("%s", err)
		os.Exit(1)
	}
}

func run() error {
	srv := api.NewServer()

	log.Info("Server listening on port :8000")
	err := http.ListenAndServe(":8000", srv)

	return err
}
