package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/avelinoschz/yofio/internal/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

// conf contains and validates the presence of environment variables
type conf struct {
	API struct {
		Host string `envconfig:"default=localhost"`
		Port string `envconfig:"default=:8000"`
	}
	Engine struct {
		Host     string `envconfig:"default=localhost"`
		Port     string `envconfig:"default=3306"`
		Dialect  string
		Database struct {
			User     string
			Password string
			Name     string
		}
	}
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	//log.SetReportCaller(true)
}

func main() {
	if err := run(); err != nil {
		log.Errorln("%s", err)
		os.Exit(1)
	}
}

func run() error {

	var conf conf
	if err := envconfig.Init(&conf); err != nil {
		return err
	}

	db, err := gorm.Open(conf.Engine.Dialect,
		fmt.Sprintf("%s:%s@/%s", conf.Engine.Database.User, conf.Engine.Database.Password, conf.Engine.Database.Name))
	defer db.Close()
	if err != nil {
		return err
	}

	srv, err := api.NewServer(db)
	if err != nil {
		return err
	}

	log.Info("Server listening on port :8000")
	err = http.ListenAndServe(conf.API.Port, srv)
	if err != nil {
		return err
	}

	return nil
}
