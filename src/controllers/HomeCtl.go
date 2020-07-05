package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*HomeCtl is a controller for home or landing page of application*/
type HomeCtl struct {
}

/*HomePage is a controller function to render home or landing page on GET request*/
func (hCtl HomeCtl) HomePage(c *gin.Context) {
	// c.String(200, "Hello, welcome the from go gin web framework v1.6.3")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go Gin Home Page",
	})
}
