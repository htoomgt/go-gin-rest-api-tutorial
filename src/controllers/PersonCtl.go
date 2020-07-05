package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoomgt/go-gin-rest-api-tutorial/src/models"
)

/*PersonCtl is a controller for person entity*/
type PersonCtl struct {
}

/*GetAll is a controller function to list down all person in json*/
func (pCtl PersonCtl) GetAll(c *gin.Context) {
	var persons []models.Person
	p := models.Person{}

	persons, err := p.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": persons,
		"count":  len(persons),
	})
}
