package model

import "database/sql"

type Person struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Surname     string         `json:"surname"`
	Patronymic  sql.NullString `json:"patronymic,omitempty"`
	Age         int            `json:"age"`
	Gender      string         `json:"gender"`
	Nationality string         `json:"nationality"`
}
