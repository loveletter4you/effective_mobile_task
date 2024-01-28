package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/loveletter4you/effective_mobile_task/internal/model"
	"github.com/loveletter4you/effective_mobile_task/internal/utils"
	"net/http"
	"strconv"
)

type personCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

func (ctr *Controller) CreatePerson(c *gin.Context) {
	var personRequest personCreateRequest
	err := c.BindJSON(&personRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person := model.Person{
		Name:    personRequest.Name,
		Surname: personRequest.Surname,
		Patronymic: sql.NullString{
			String: personRequest.Patronymic,
			Valid:  true,
		},
	}
	if personRequest.Patronymic == "" {
		person.Patronymic.Valid = false
	}
	person.Age, err = utils.GetAge(person.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person.Gender, err = utils.GetGender(person.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person.Nationality, err = utils.GetNationality(person.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctr.storage.Person().InsertPerson(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (ctr *Controller) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := ctr.storage.Person().DeletePerson(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (ctr *Controller) GetPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := ctr.storage.Person().GetPerson(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (ctr *Controller) GetPersons(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	if page < 0 {
		page = 0
	}
	if limit < 1 {
		limit = 1
	} else if limit > 100 {
		limit = 100
	}
	persons, err := ctr.storage.Person().GetPersons(limit, page)
	c.JSON(http.StatusOK, persons)
}
