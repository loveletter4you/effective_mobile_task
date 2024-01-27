package app

import (
	"github.com/loveletter4you/effective_mobile_task/config"
	"log"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cfg)
}
