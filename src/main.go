package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/htoomgt/go-gin-rest-api-tutorial/src/routes"

	"github.com/htoomgt/go-gin-rest-api-tutorial/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/htoomgt/go-gin-rest-api-tutorial/src/models"
)

/*HomePage is a controller function
with GET Method to render or result home page
*/
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World !",
	})
}

/*PostHomePage is a controller function
with POST method to render or result home page*/
func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("There is error ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

/*QueryStrings is a controller function
with GET method to render value from url query param*/
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

/*PathParamters is a controller function
with GET method to render value from url path param*/
func PathParamters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

/*PersonGetByID is a function to get person data by ID*/
func PersonGetByID(c *gin.Context) {
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

/*StorePerson is a controller function wit GET method
to store a person into mysql database*/
func StorePerson(c *gin.Context) {
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

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)            // /query?name=rick&age=32
	r.GET("/path/:name/:age", PathParamters) // /path/rick/32

	var personCtl controllers.PersonCtl
	r.POST("/person", StorePerson)
	r.GET("/persons", personCtl.GetAll)
	r.GET("/person/:id", PersonGetByID)

	r.GET("/getCurrentTime", func(c *gin.Context) {
		dt := time.Now()
		nowDt := dt.Format("2006-01-02 15:04:05")

		fmt.Println("Current Date Time is ")

		c.JSON(http.StatusOK, gin.H{
			"current date time is ": nowDt,
		})
	})

	var apiRoutes routes.APIRoutes
	apiRoutes.RegisteredRoutes(r)

	r.Run(":8181")
}
