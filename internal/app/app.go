package app

import (
	"github.com/loveletter4you/effective_mobile_task/config"
	"github.com/loveletter4you/effective_mobile_task/internal/router"
	"log"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	r := router.NewServer()
	err = r.StartServer(cfg)
	log.Fatal(err)
}
