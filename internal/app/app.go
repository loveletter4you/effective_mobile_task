package app

import (
	"github.com/loveletter4you/effective_mobile_task/config"
	"github.com/loveletter4you/effective_mobile_task/internal/storage"
	"github.com/loveletter4you/effective_mobile_task/internal/utils"
	"log"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cfg)
	s := storage.NewStorage()
	if err := s.Open(cfg); err != nil {
		log.Fatal(err)
	}
	log.Print(s)
	age, err := utils.GetAge("Dmitry")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(age)
	gender, err := utils.GetGender("Dmitry")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(gender)
	nationality, err := utils.GetNationality("Dmitry")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(nationality)
}
