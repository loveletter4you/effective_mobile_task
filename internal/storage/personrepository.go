package storage

import (
	"fmt"
	"github.com/loveletter4you/effective_mobile_task/internal/model"
)

type PersonRepository struct {
	storage *Storage
}

func (pr *PersonRepository) InsertPerson(person *model.Person) error {
	var query string
	if person.Patronymic != "" {
		query = fmt.Sprintf("INSERT INTO persons (name, surname, patronymic, age, nationality, gender) "+
			"VALUES ($$%s$$, $$%s$$, $$%s$$, %d, $$%s$$, $$%s$$) RETURNING id",
			person.Name, person.Surname, person.Patronymic, person.Age, person.Nationality, person.Gender)
	} else {
		query = fmt.Sprintf("INSERT INTO persons (name, surname, age, nationality, gender) "+
			"VALUES ($$%s$$, $$%s$$, %d, $$%s$$, $$%s$$) RETURNING id",
			person.Name, person.Surname, person.Age, person.Nationality, person.Gender)
	}
	return pr.storage.db.QueryRow(query).Scan(&person.Id)
}

func (pr *PersonRepository) DeletePerson(id int) error {
	query := fmt.Sprintf("DELETE FROM persons WHERE id = %d", id)
	return pr.storage.db.QueryRow(query).Err()
}

func (pr *PersonRepository) GetPerson(id int) (*model.Person, error) {
	query := fmt.Sprintf("SELECT * FROM persons WHERE id = %d", id)
	var person *model.Person
	err := pr.storage.db.QueryRow(query).Scan(&person)
	return person, err
}

func (pr *PersonRepository) GetPersons(limit, page int) ([]*model.Person, error) {
	query := fmt.Sprintf("SELECT * FROM persons OFFSET %d LIMIT %d", limit*page, limit)
	rows, err := pr.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := make([]*model.Person, limit)
	for rows.Next() {
		person := &model.Person{}
		if err := rows.Scan(&person); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}
