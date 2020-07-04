package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

/*StorePerson is a controller function wit GET method
to store a person into mysql database*/
func StorePerson(c *gin.Context) {
	//body := c.Request.Body

}

/*AllPerson is a controller function wit GET method
to get all person into mysql database*/
func AllPerson(c *gin.Context) {
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

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)            // /query?name=rick&age=32
	r.GET("/path/:name/:age", PathParamters) // /path/rick/32

	r.GET("/store-person", StorePerson)
	r.GET("/get-all-person", AllPerson)

	r.Run(":8181")
}
