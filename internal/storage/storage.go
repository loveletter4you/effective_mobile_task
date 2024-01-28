package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/loveletter4you/effective_mobile_task/config"
	"time"
)

type Storage struct {
	db               *sql.DB
	personRepository *PersonRepository
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Open(config *config.Config) error {
	var (
		db  *sql.DB
		err error
	)

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Database.Username,
		config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	for config.DatabaseConnection.Attempts > 0 {
		db, err = sql.Open("postgres", dbUrl)
		if err != nil {
			break
		}

		if err = db.Ping(); err == nil {
			s.db = db
			break
		}
		time.Sleep(time.Duration(config.DatabaseConnection.Timeout) * time.Second)
		config.DatabaseConnection.Attempts--
	}
	return err
}

func (s *Storage) Close() error {
	err := s.db.Close()
	return err
}

func (s *Storage) Person() *PersonRepository {
	if s.personRepository != nil {
		return s.personRepository
	}

	s.personRepository = &PersonRepository{
		storage: s,
	}

	return s.personRepository
}
