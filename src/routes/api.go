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

	/*api for setting /api uri*/
	api := router.Group("/api")
	{
		/*v1 under api for setting /api/v1 uri*/
		v1 := api.Group("/v1")
		{
			/*Testing route from route file*/
			v1.GET("/say-hello", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "hello world!",
				})
			})

			/*Person related routes block start*/
			v1.GET("/persons", personCtl.GetAll)
			v1.GET("/person/:id", personCtl.GetByID)
			v1.POST("/person", personCtl.Add)
			v1.PUT("/person/:id", personCtl.UpdateAllByID)
			v1.DELETE("/person/:id", personCtl.DeleteByID)
			v1.GET("/faker-persons-add/:count", personCtl.AddFakerGenPersons)
			/*Person related routes block end*/
		}
	}
}
