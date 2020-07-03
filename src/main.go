package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
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

	c.JSON(200, gin.H{
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

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)            // /query?name=rick&age=32
	r.GET("/path/:name/:age", PathParamters) // /path/rick/32

	r.Run(":8181")
}
