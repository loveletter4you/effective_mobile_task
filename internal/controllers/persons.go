package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/loveletter4you/effective_mobile_task/internal/model"
	"github.com/loveletter4you/effective_mobile_task/internal/utils"
	"net/http"
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
		Name:       personRequest.Name,
		Surname:    personRequest.Surname,
		Patronymic: personRequest.Patronymic,
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
