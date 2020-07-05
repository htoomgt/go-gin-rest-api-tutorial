package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

/*GetByID is a controller function to get a person by ID*/
func (pCtl PersonCtl) GetByID(c *gin.Context) {
	var result gin.H
	id := c.Param("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	p := models.Person{
		ID: ID,
	}
	person, err := p.GetByID()
	if err != nil {
		result = gin.H{
			"count":  0,
			"result": nil,
		}
	} else {
		result = gin.H{
			"count":  1,
			"result": person,
		}
	}
	c.JSON(http.StatusOK, result)
}

/*Add is a controller function to add a new person*/
func (pCtl PersonCtl) Add(c *gin.Context) {
	var p models.Person
	err := c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	ID, err := p.Add()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(ID)
	name := p.FirstName + " " + p.LastName
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name),
	})
}

/*UpdateAllByID is a controller function to update all fields by ID*/
func (pCtl PersonCtl) UpdateAllByID(c *gin.Context) {
	var (
		p      models.Person
		buffer bytes.Buffer
	)

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	p.ID = ID
	rows, err := p.UpdateAllByID()
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(rows)
	buffer.WriteString(p.FirstName)
	buffer.WriteString(" ")
	buffer.WriteString(p.LastName)
	name := buffer.String()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated to %s . Affected row count %d", name, rows),
	})

}

/*DeleteByID is a controller function to delete person by ID*/
func (pCtl PersonCtl) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}
	p := models.Person{
		ID: ID,
	}

	err = c.Bind(&p)
	if err != nil {
		log.Fatalln(err)
	}

	rowCount, err := p.DeleteByID()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted with person ID %d. Affected row count %d", ID, rowCount),
	})
}

/*AddFakerGenPersons is a controller function to add faker generated persons
to database as specificed number to check the performace and to understand or
 test non-blocking feature of golang.*/
func (pCtl PersonCtl) AddFakerGenPersons(c *gin.Context) {

}
