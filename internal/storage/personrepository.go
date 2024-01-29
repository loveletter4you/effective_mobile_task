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
	if person.Patronymic.Valid == true {
		query = fmt.Sprintf("INSERT INTO persons (name, surname, patronymic, age, nationality, gender) "+
			"VALUES ($$%s$$, $$%s$$, $$%s$$, %d, $$%s$$, $$%s$$) RETURNING id",
			person.Name, person.Surname, person.Patronymic.String, person.Age, person.Nationality, person.Gender)
	} else {
		query = fmt.Sprintf("INSERT INTO persons (name, surname, age, nationality, gender) "+
			"VALUES ($$%s$$, $$%s$$, %d, $$%s$$, $$%s$$) RETURNING id",
			person.Name, person.Surname, person.Age, person.Nationality, person.Gender)
	}
	return pr.storage.db.QueryRow(query).Scan(&person.Id)
}

func (pr *PersonRepository) UpdatePerson(person *model.Person) error {
	var query string
	if person.Patronymic.Valid == true {
		query = fmt.Sprintf("UPDATE persons "+
			"SET name = $$%s$$, surname = $$%s$$, patronymic = $$%s$$, age = %d,nationality = $$%s$$, gender = $$%s$$ "+
			"WHERE id = %d", person.Name, person.Surname, person.Patronymic.String,
			person.Age, person.Nationality, person.Gender, person.Id)
	} else {
		query = fmt.Sprintf("UPDATE persons "+
			"SET name = $$%s$$, surname = $$%s$$, age = %d,nationality = $$%s$$, gender = $$%s$$ "+
			"WHERE id = %d", person.Name, person.Surname,
			person.Age, person.Nationality, person.Gender, person.Id)
	}
	return pr.storage.db.QueryRow(query).Err()
}

func (pr *PersonRepository) DeletePerson(id int) (*model.Person, error) {
	query := fmt.Sprintf("DELETE FROM persons WHERE id = %d RETURNING "+
		"id, name, surname, patronymic, age, nationality, gender", id)
	person := &model.Person{}
	err := pr.storage.db.QueryRow(query).Scan(&person.Id, &person.Name, &person.Surname,
		&person.Patronymic, &person.Age, &person.Nationality, &person.Gender)
	return person, err
}

func (pr *PersonRepository) GetPerson(id int) (*model.Person, error) {
	query := fmt.Sprintf("SELECT id, name, surname, patronymic, age, nationality, gender FROM persons WHERE id = %d", id)
	person := &model.Person{}
	err := pr.storage.db.QueryRow(query).Scan(&person.Id, &person.Name, &person.Surname,
		&person.Patronymic, &person.Age, &person.Nationality, &person.Gender)
	return person, err
}

func (pr *PersonRepository) GetPersons(limit, page int) ([]*model.Person, error) {
	query := fmt.Sprintf("SELECT id, name, surname, patronymic, age, nationality, gender "+
		"FROM persons OFFSET %d LIMIT %d", limit*page, limit)
	rows, err := pr.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := make([]*model.Person, 0)
	for rows.Next() {
		person := &model.Person{}
		if err := rows.Scan(&person.Id, &person.Name, &person.Surname,
			&person.Patronymic, &person.Age, &person.Nationality, &person.Gender); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}
