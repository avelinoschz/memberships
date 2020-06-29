package main

import (
	"net/http"
	"os"

	"github.com/avelinoschz/yofio/internal/api"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

var conf struct {
	API struct {
		Host string
		Port string
	}
}

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

	if err := envconfig.Init(&conf); err != nil {
		return err
	}

	srv := api.NewServer()

	log.Info("Server listening on port :8000")
	err := http.ListenAndServe(":8000", srv)
	if err != nil {
		return err
	}

	return nil
}
