package main

import (
	"net/http"
	"os"

	"github.com/avelinoschz/yofio/internal/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return err
	}

	srv := api.NewServer(db)

	log.Info("Server listening on port :8000")
	err = http.ListenAndServe(":8000", srv)
	if err != nil {
		return err
	}

	return nil
}
