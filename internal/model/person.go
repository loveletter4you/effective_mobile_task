package model

import "database/sql"

type Person struct {
	Id          int            `json:"id"`
	Name        string         `json:"name" binding:"required"`
	Surname     string         `json:"surname" binding:"required"`
	Patronymic  sql.NullString `json:"patronymic,omitempty"`
	Age         int            `json:"age" binding:"required"`
	Gender      string         `json:"gender" binding:"required"`
	Nationality string         `json:"nationality" binding:"required"`
}
