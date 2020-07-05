package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/htoomgt/go-gin-rest-api-tutorial/src/controllers"
)

/*WebRoutes will be used for web route separations*/
type WebRoutes struct {
}

/*RegisteredRoutes is a function to register routes*/
func (webRoutes WebRoutes) RegisteredRoutes(router *gin.Engine) {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	router.LoadHTMLFiles(rootDir + "/src/resources/views/index.tmpl")

	var homeCtl controllers.HomeCtl
	router.GET("/", homeCtl.HomePage)
}
