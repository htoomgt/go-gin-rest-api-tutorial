package routes

import (
	"net/http"

	"github.com/htoomgt/go-gin-rest-api-tutorial/src/controllers"

	"github.com/gin-gonic/gin"
)

/*APIRoutes will be used for route separations*/
type APIRoutes struct {
}

/*RegisteredRoutes is a function to register routes*/
func (apiRoutes APIRoutes) RegisteredRoutes(router *gin.Engine) {
	var personCtl controllers.PersonCtl

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/say-hello", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "hello world!",
				})
			})

			v1.GET("/persons", personCtl.GetAll)
			// v1.GET("//person/:id", per)
		}
	}
}
